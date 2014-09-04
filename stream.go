package nest

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strings"
)

/*
DevicesStream emits events from the Nest devices REST streaming API

	client.DevicesStream(func(event *Devices) {
		fmt.Println(event)
	})
*/
func (c *Client) DevicesStream(callback func(devices *Devices, err error)) {
	c.setRedirectURL()
	for {
		c.streamDevices(callback)
	}
}

// streamDevices connects to the stream, following the redirect and then watches the stream
func (c *Client) streamDevices(callback func(devices *Devices, err error)) {
	resp, err := c.getDevices(Stream)
	if err != nil {
		callback(nil, err)
		return
	}
	defer resp.Body.Close()
	c.watchDevicesStream(resp, callback)
}

// watchDevicesStream grabs the data off the stream, parses them and invokes the callback
func (c *Client) watchDevicesStream(resp *http.Response, callback func(devices *Devices, err error)) {
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		value := parseStreamData(line)
		if value != "" {
			event := &Event{}
			json.Unmarshal([]byte(value), event)
			if event.Data != nil {
				c.associateClientToDevices(event.Data)
				callback(event.Data, nil)
			}
		}
	}
}

// parseStreamData takes a line of the stream and parses out the JSON data
func parseStreamData(line string) string {
	sections := strings.SplitN(line, ":", 2)
	field, value := sections[0], ""
	if len(sections) == 2 {
		value = strings.TrimPrefix(sections[1], " ")
	}
	if field == "data" {
		return value
	}
	return ""
}
