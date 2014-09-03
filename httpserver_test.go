package nest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func serveHTTP(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()
		switch req.RequestURI {
		case "/?code=ABCD1234&client_id=1234&client_secret=5678&grant_type=authorization_code":
			w.WriteHeader(400)
			w.Write(oauthErrorJSON())
		case "/?code=EFGH5678&client_id=1234&client_secret=5678&grant_type=authorization_code":
			w.WriteHeader(200)
			w.Write(successTokenJSON())
		case "/devices.json?auth=" + Token:
			if req.Header.Get("Accept") == "text/event-stream" {
				f, _ := w.(http.Flusher)
				w.Header().Set("Content-Type", "text/event-stream")
				w.Header().Set("Cache-Control", "no-cache")
				w.Header().Set("Connection", "keep-alive")
				fmt.Fprintf(w, "data: %s\n\n", streamEvent())
				f.Flush()
				fmt.Fprintf(w, "data: %s\n\n", streamEvent())
				f.Flush()
			} else {
				urlStr := "/redirected/" + req.URL.String()
				http.Redirect(w, req, urlStr, 307)
			}
		case "/redirected/devices.json?auth=" + Token:
			w.WriteHeader(200)
			w.Write(devicesResponseJSON())
		case "/devices/thermostats/z1234?auth=" + Token:
			if strings.Contains(string(body), "fan_timer_active") {
				w.WriteHeader(200)
				w.Write(body)
				return
			}
			if strings.Contains(string(body), "hvac_mode") {
				w.WriteHeader(200)
				w.Write(body)
				return
			}
			if strings.Contains(string(body), "600") {
				w.WriteHeader(400)
				w.Write([]byte("Bad Request"))
				return
			}
			if strings.Contains(string(body), "high") {
				w.WriteHeader(200)
				w.Write(body)
				return
			}
			w.WriteHeader(200)
			if strings.Contains(string(body), "target_temperature_c") {
				w.Write([]byte(`{"target_temperature_c":28.5}`))
			} else {
				w.Write([]byte(`{"target_temperature_f":50}`))
			}
		}
	}))
}

func oauthErrorJSON() []byte {
	return []byte(`
		{
		    "error": "oauth2_error",
		    "error_description": "authorization code not found"
		}
		`)
}

func successTokenJSON() []byte {
	return []byte(`
		{
		    "access_token": "` + Token + `",
		    "expires_in": 315360000
		}
		`)
}

func devicesResponseJSON() []byte {
	return []byte(`
		{
		    "thermostats": {
		        "z1234": {
		            "locale": "en-US",
		            "temperature_scale": "F",
		            "is_using_emergency_heat": false,
		            "has_fan": false,
		            "software_version": "4.2.4",
		            "has_leaf": true,
		            "device_id": "z1234",
		            "name": "Bedroom (Main)",
		            "can_heat": true,
		            "can_cool": false,
		            "hvac_mode": "heat",
		            "target_temperature_c": 10,
		            "target_temperature_f": 50,
		            "target_temperature_high_c": 24,
		            "target_temperature_high_f": 75,
		            "target_temperature_low_c": 20,
		            "target_temperature_low_f": 68,
		            "ambient_temperature_c": 21.5,
		            "ambient_temperature_f": 72,
		            "away_temperature_high_c": 24,
		            "away_temperature_high_f": 76,
		            "away_temperature_low_c": 10,
		            "away_temperature_low_f": 50,
		            "structure_id": "s1234",
		            "fan_timer_active": false,
		            "name_long": "Bedroom Thermostat (Main)",
		            "is_online": true,
		            "last_connection": "2014-08-28T23:03:03.439Z"
		        }
		    },
		    "smoke_co_alarms": {
		        "z5678": {
		            "name": "Upstairs Hallway",
		            "locale": "en-US",
		            "structure_id": "s1234",
		            "software_version": "1.0rc12",
		            "device_id": "z5678",
		            "name_long": "Upstairs Hallway Nest Protect",
		            "is_online": true,
		            "last_connection": "2014-08-28T07:35:46.542Z",
		            "battery_health": "ok",
		            "co_alarm_state": "ok",
		            "smoke_alarm_state": "ok",
		            "ui_color_state": "green"
		        },
		        "z90123": {
		            "name": "Downstairs Hallway",
		            "locale": "en-US",
		            "structure_id": "s1234",
		            "software_version": "1.0rc12",
		            "device_id": "z90123",
		            "name_long": "Downstairs Hallway Nest Protect",
		            "is_online": true,
		            "last_connection": "2014-08-28T07:08:17.390Z",
		            "battery_health": "ok",
		            "co_alarm_state": "ok",
		            "smoke_alarm_state": "ok",
		            "ui_color_state": "green"
		        }
		    }
		}
		`)
}

func streamEvent() []byte {
	return []byte(`{"path":"/devices","data":{"thermostats":{"z1234":{"locale":"en-US","temperature_scale":"F","is_using_emergency_heat":false,"has_fan":true,"software_version":"4.1","has_leaf":true,"device_id":"z1234","name":"Entryway","can_heat":true,"can_cool":true,"hvac_mode":"heat","target_temperature_c":19.0,"target_temperature_f":67,"target_temperature_high_c":24.0,"target_temperature_high_f":75,"target_temperature_low_c":20.0,"target_temperature_low_f":68,"ambient_temperature_c":21.0,"ambient_temperature_f":70,"away_temperature_high_c":24.0,"away_temperature_high_f":76,"away_temperature_low_c":12.5,"away_temperature_low_f":55,"structure_id":"s1234","fan_timer_active":false,"name_long":"Entryway Thermostat","is_online":true},"z5678":{"locale":"en-US","temperature_scale":"F","is_using_emergency_heat":false,"has_fan":false,"software_version":"4.2.4","has_leaf":true,"device_id":"Zz5678","name":"Bedroom (Master)","can_heat":true,"can_cool":false,"hvac_mode":"heat","target_temperature_c":10.0,"target_temperature_f":50,"target_temperature_high_c":24.0,"target_temperature_high_f":75,"target_temperature_low_c":20.0,"target_temperature_low_f":68,"ambient_temperature_c":20.5,"ambient_temperature_f":70,"away_temperature_high_c":24.0,"away_temperature_high_f":76,"away_temperature_low_c":10.0,"away_temperature_low_f":50,"structure_id":"s1234","fan_timer_active":false,"name_long":"Bedroom Thermostat (Master)","is_online":true,"last_connection":"2014-08-30T16:29:45.165Z"}},"smoke_co_alarms":{"a1234":{"name":"Upstairs Hallway","locale":"en-US","structure_id":"a1234","software_version":"1.0rc12","device_id":"a1234","name_long":"Upstairs Hallway Nest Protect","is_online":true,"last_connection":"2014-08-30T05:35:47.025Z","battery_health":"ok","co_alarm_state":"ok","smoke_alarm_state":"ok","ui_color_state":"green"},"a3455":{"name":"Bedroom","locale":"en-US","structure_id":"s1234","software_version":"1.0.2rc2","device_id":"a3455","name_long":"Bedroom Nest Protect","is_online":true,"battery_health":"ok","co_alarm_state":"ok","smoke_alarm_state":"ok","ui_color_state":"green"},"a6789":{"name":"Downstairs Hallway","locale":"en-US","structure_id":"s1234","software_version":"1.0rc12","device_id":"a6789","name_long":"Downstairs Hallway Nest Protect","is_online":true,"last_connection":"2014-08-30T05:08:17.377Z","battery_health":"ok","co_alarm_state":"ok","smoke_alarm_state":"ok","ui_color_state":"green"}}}}`)
}
