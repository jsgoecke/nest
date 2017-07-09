package nest

import (
	"encoding/json"
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

/*
Authorize fetches and sets the Nest API token
https://developer.nest.com/documentation/how-to-auth

	client.Authorize()
*/
func (c *Client) Authorize() *APIError {
	resp, err := http.Post(c.authURL(), "application/x-www-form-urlencoded", nil)
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
	if c.RedirectURL == "" {
		req, _ := http.NewRequest("GET", c.APIURL+"/devices.json?auth="+c.Token, nil)
		resp, err := http.DefaultClient.Do(req)
		if resp.Request.URL != nil {
			c.RedirectURL = resp.Request.URL.Scheme + "://" + resp.Request.URL.Host
		}
		return resp, err
	}

	req, _ := http.NewRequest("GET", c.RedirectURL+"/devices.json?auth="+c.Token, nil)
	if action == Stream {
		req.Header.Set("Accept", "text/event-stream")
	}
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}

// authURL sets the full authorization URL for the Nest API
func (c *Client) authURL() string {
	location := c.AccessTokenURL + "?code=" + c.AuthorizationCode
	location += "&client_id=" + c.ID
	location += "&client_secret=" + c.Secret
	location += "&grant_type=authorization_code"
	return location
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

// setRedirectURL sets the URL if not already set
func (c *Client) setRedirectURL() (int, error) {
	if c.RedirectURL == "" {
		resp, err := c.getDevices(NoStream)
		if err != nil || resp.StatusCode != 200 {
			return resp.StatusCode, err
		}
	}
	return 0, nil
}
