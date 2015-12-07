package nest

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
	"time"
)

func TestCombined(t *testing.T) {
	Convey("Given a JSON object with combined", t, func() {
		combined := &Combined{}
		err := json.Unmarshal(combinedJSON(), combined)
		So(err, ShouldBeNil)

		Convey("We should get thermostats", func() {
			So(len(combined.Devices.Thermostats), ShouldEqual, 1)
			for key, value := range combined.Devices.Thermostats {
				So(key, ShouldEqual, combined.Devices.Thermostats["peyiJNo0IldT2YlIVtYaGQ"].DeviceID)
				checkFields(value)
			}
		})

		Convey("We should get smokecoalarms", func() {
			So(len(combined.Devices.SmokeCoAlarms), ShouldEqual, 1)
			for key, value := range combined.Devices.SmokeCoAlarms {
				So(key, ShouldEqual, combined.Devices.SmokeCoAlarms["RTMTKxsQTCxzVcsySOHPxKoF4OyCifrs"].DeviceID)
				checkFields(value)
			}
		})

		Convey("We should get structures", func() {
			So(len(combined.Structures), ShouldEqual, 1)
			for key, value := range combined.Structures {
				So(key, ShouldEqual, combined.Structures["VqFabWH21nwVyd4RWgJgNb292wa7hG_dUwo2i2SG7j3-BOLY0BA4sw"].StructureID)
				checkFields(value)
			}
		})

		Convey("Should get an eta", func() {
			So(combined.Structures["VqFabWH21nwVyd4RWgJgNb292wa7hG_dUwo2i2SG7j3-BOLY0BA4sw"].ETA.TripID, ShouldEqual, "myTripHome1024")
			checkFields(combined.Structures["VqFabWH21nwVyd4RWgJgNb292wa7hG_dUwo2i2SG7j3-BOLY0BA4sw"].ETA)
		})
	})
}

func TestAccess(t *testing.T) {
	Convey("Given a JSON object with an access token", t, func() {
		access := &Access{}
		err := json.Unmarshal(accessJSON(), access)
		So(err, ShouldBeNil)
	})
}

func checkFields(value interface{}) {
	s := reflect.ValueOf(value).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		value := f.Interface()
		switch value.(type) {
		case string:
			So(value.(string), ShouldNotBeEmpty)
		case bool:
			So(value.(bool), ShouldBeTrue)
		case int:
			So(value.(int), ShouldNotBeNil)
		case float32:
			So(value.(float32), ShouldNotBeNil)
		case time.Time:
			So(value.(time.Time), ShouldNotBeNil)
		}
	}
}

func combinedJSON() []byte {
	return []byte(`
		{
		    "devices": {
		        "thermostats": {
		            "peyiJNo0IldT2YlIVtYaGQ": {
		                "device_id": "peyiJNo0IldT2YlIVtYaGQ",
		                "locale": "en-US",
		                "software_version": "4.0",
		                "structure_id": "VqFabWH21nwVyd4RWgJgNb292wa7hG_dUwo2i2SG7j3-BOLY0BA4sw",
		                "name": "Hcombinedway (upstairs)",
		                "name_long": "Hcombinedway Thermostat (upstairs)",
		                "last_connection": "2014-03-02T23:20:19+00:00",
		                "is_online": true,
		                "can_cool": true,
		                "can_heat": true,
		                "is_using_emergency_heat": true,
		                "has_fan": true,
		                "fan_timer_active": true,
		                "fan_timer_timeout": "2014-03-02T23:20:19+00:00",
		                "has_leaf": true,
		                "temperature_scale": "C",
		                "target_temperature_f": 72,
		                "target_temperature_c": 21.5,
		                "target_temperature_high_f": 72,
		                "target_temperature_high_c": 21.5,
		                "target_temperature_low_f": 64,
		                "target_temperature_low_c": 17.5,
		                "away_temperature_high_f": 72,
		                "away_temperature_high_c": 21.5,
		                "away_temperature_low_f": 64,
		                "away_temperature_low_c": 17.5,
		                "hvac_mode": "heat",
		                "ambient_temperature_f": 72,
		                "ambient_temperature_c": 21.5,
		                "humidity": 35,
		                "hvac_state": "heating",
		                "where_id": "d6reb_OZTM..."
		            }
		        },
		        "smoke_co_alarms": {
		            "RTMTKxsQTCxzVcsySOHPxKoF4OyCifrs": {
		                "device_id": "RTMTKxsQTCxzVcsySOHPxKoF4OyCifrs",
		                "locale": "en-US",
		                "software_version": "1.01",
		                "structure_id": "VqFabWH21nwVyd4RWgJgNb292wa7hG_dUwo2i2SG7j3-BOLY0BA4sw",
		                "name": "Hcombinedway (upstairs)",
		                "name_long": "Hcombinedway Protect (upstairs)",
		                "last_connection": "2014-03-02T23:20:19+00:00",
		                "is_online": true,
		                "battery_health": "ok",
		                "co_alarm_state": "ok",
		                "smoke_alarm_state": "ok",
		                "ui_color_state": "gray"
		            }
		        }
		    },
		    "structures": {
		        "VqFabWH21nwVyd4RWgJgNb292wa7hG_dUwo2i2SG7j3-BOLY0BA4sw": {
		            "structure_id": "VqFabWH21nwVyd4RWgJgNb292wa7hG_dUwo2i2SG7j3-BOLY0BA4sw",
		            "thermostats": [
		                "peyiJNo0IldT2YlIVtYaGQ"
		            ],
		            "smoke_co_alarms": [
		                "RTMTKxsQTCxzVcsySOHPxKoF4OyCifrs"
		            ],
		            "away": "home",
		            "name": "Home",
		            "country_code": "US",
		            "peak_period_start_time": "2014-03-10T23:10:12+00:00",
		            "peak_period_end_time": "2014-03-10T23:14:19+00:00",
		            "time_zone": "America/Los_Angeles",
		            "eta": {
		                "trip_id": "myTripHome1024",
		                "estimated_arrival_window_begin": "2014-07-04T10:48:11+00:00",
		                "estimated_arrival_window_end": "2014-07-04T18:48:11+00:00"
		            }
		        }
		    }
		}`)
}

func accessJSON() []byte {
	return []byte(`
		{
			"access_token": "c.FmDPkzyzaQe...",
			"expires_in": 315360000
		}`)
}
