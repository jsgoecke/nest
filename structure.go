package nest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
Structures returns a map of structures
https://developer.nest.com/documentation/api#structures

	structures := client.Structures()
*/
func (c *Client) Structures() (map[string]*Structure, *APIError) {
	resp, err := c.getStructures()
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
	structures := make(map[string]*Structure)
	err = json.Unmarshal(body, &structures)
	c.associateClientToStructures(structures)
	return structures, nil
}

/*
SetAway sets the away status of a structure
https://developer.nest.com/documentation/api#away

	s.SetAway(nest.Away)
*/
func (s *Structure) SetAway(mode int) *APIError {
	requestMode := make(map[string]string)
	switch mode {
	case Home:
		requestMode["away"] = "home"
	case Away:
		requestMode["away"] = "away"
	case AutoAway:
		requestMode["away"] = "auto-away"
	default:
		return generateAPIError("Invalid Away requested - must be home, away or auto-away")
	}
	body, _ := json.Marshal(requestMode)
	return s.setStructure(body)
}

// getStructures does an HTTP get
func (c *Client) getStructures() (*http.Response, error) {
	if c.RedirectURL == "" {
		req, _ := http.NewRequest("GET", c.APIURL+"/structures.json?auth="+c.Token, nil)
		resp, err := http.DefaultClient.Do(req)
		if resp.Request.URL != nil {
			c.RedirectURL = resp.Request.URL.Scheme + "://" + resp.Request.URL.Host
		}
		return resp, err
	}

	req, _ := http.NewRequest("GET", c.RedirectURL+"/structures.json?auth="+c.Token, nil)
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}

// setStructure sends the request to the Nest REST API
func (s *Structure) setStructure(body []byte) *APIError {
	url := s.Client.RedirectURL + "/structures/" + s.StructureID + "?auth=" + s.Client.Token
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		apiError := &APIError{
			Error:       "http_error",
			Description: err.Error(),
		}
		return apiError
	}
	body, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		structure := &Structure{}
		json.Unmarshal(body, structure)
		return nil
	}
	apiError := &APIError{}
	json.Unmarshal(body, apiError)
	apiError = generateAPIError(apiError.Error)
	apiError.Status = resp.Status
	apiError.StatusCode = resp.StatusCode
	return apiError
}

// associateClientToStructures ensures each structure knows its client details
func (c *Client) associateClientToStructures(structures map[string]*Structure) {
	for _, value := range structures {
		value.Client = c
	}
}
