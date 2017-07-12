package nest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	// Version is the version of the library
	Version = "0.0.1"
	// APIURL is the main URL for the API
	APIURL = "https://developer-api.nest.com"
	// AccessTokenURL is the Next API URL to get an access_token
	AccessTokenURL = "https://api.home.nest.com/oauth2/access_token"
	// NoStream indicates wedo not want to stream on a GET for server side events
	NoStream = iota
	//Stream indicates we want to stream on a GET for server side events
	Stream
	// Cool sets HvacMode to "cool"
	Cool
	// Heat sets HvacMode to "heat"
	Heat
	// HeatCool sets HvacMode to "heat-cool"
	HeatCool
	// Eco sets the HVacMode to "eco"
	Eco
	// Off sets HvacMode to "off"
	Off
	// Home sets Away mode to "home"
	Home
	// Away sets Away mode to "away"
	Away
	// AutoAway sets Away mode to "auto-away"
	AutoAway
)

/* Configure a httpClient that will handle redirects */
var httpClient = &http.Client{
	CheckRedirect: func(redirRequest *http.Request, via []*http.Request) error {
		// Go's http.DefaultClient does not forward headers when a redirect 3xx
		// response is recieved. Thus, the header (which in this case contains the
		// Authorization token) needs to be passed forward to the redirect
		// destinations.
		redirRequest.Header = via[0].Header

		// Go's http.DefaultClient allows 10 redirects before returning an
		// an error. We have mimicked this default behavior.s
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}

		return nil
	},
}

/*
New creates a new Nest client

	client := New("1234", "STATE", "<secret>", "<auth-code>")
*/
func New(clientID string, state string, clientSecret string, authorizationCode string) *Client {
	return &Client{
		ClientID:          clientID,
		State:             state,
		ClientSecret:      clientSecret,
		AuthorizationCode: authorizationCode,
		AccessTokenURL:    AccessTokenURL,
		APIURL:            APIURL,
	}
}

func NewWithAuthorization(AccessToken string) *Client {
	return &Client{
		Token:  AccessToken,
		APIURL: APIURL,
	}
}

/*
Authorize fetches and sets the Nest API token
https://developer.nest.com/documentation/how-to-auth

	client.Authorize()
*/
func (c *Client) Authorize() *APIError {
	req, err := http.NewRequest(http.MethodPost, c.AccessTokenURL, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("client_id", c.ClientID)
	req.Header.Add("client_secret", c.ClientSecret)
	req.Header.Add("grant_type", "authorization_code")

	var client = &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &APIError{
			Error:       "http_error",
			Description: err.Error(),
			Status:      resp.Status,
			StatusCode:  resp.StatusCode,
		}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &APIError{
			Error:       "body_error",
			Description: err.Error(),
		}
	}
	if resp.StatusCode != 200 {
		apiError := &APIError{}
		json.Unmarshal(body, apiError)
		return apiError
	}
	access := &Access{}
	json.Unmarshal(body, access)
	c.Token = access.Token
	c.ExpiresIn = access.ExpiresIn
	return nil
}

/*
Devices returns a list of devices

	devices := client.Devices()
*/
func (c *Client) Devices() (*Devices, *APIError) {
	resp, err := c.getDevices(NoStream)
	if err != nil {
		return nil, &APIError{
			Error:       "devices_error",
			Description: err.Error(),
		}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &APIError{
			Error:       "body_read_error",
			Description: err.Error(),
		}
	}
	if resp.StatusCode != 200 {
		apiError := &APIError{}
		json.Unmarshal(body, apiError)
		return nil, apiError
	}
	devices := &Devices{}
	err = json.Unmarshal(body, devices)
	c.associateClientToDevices(devices)
	return devices, nil
}

// getDevices does an HTTP get with or without a stream on devices
func (c *Client) getDevices(action int) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.APIURL+"/devices.json", nil)
	req.Header.Add("Content-Type", "\"application/json\"")
	req.Header.Add("Authorization", c.Token)

	if action == Stream {
		req.Header.Set("Accept", "text/event-stream")
	}

	response, err := httpClient.Do(req)

	return response, err
}

// associateClientToDevices ensures each device knows its client details
func (c *Client) associateClientToDevices(devices *Devices) {
	for _, value := range devices.Thermostats {
		value.Client = c
	}
	for _, value := range devices.SmokeCoAlarms {
		value.Client = c
	}
	for _, value := range devices.Cameras {
		value.Client = c
	}
}
