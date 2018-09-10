package nest

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFanTimeActive(t *testing.T) {
	Convey("When setting the Fan Timer", t, func() {
		client := New(ClientID, State, ClientSecret, AuthorizationCode)
		client.AccessTokenURL = ts.URL
		client.Authorize()
		client.Token = Token
		client.APIURL = ts.URL
		devices, _ := client.Devices()
		Convey("When setting the Fan Timer to active", func() {
			err := devices.Thermostats["z1234"].SetFanTimerActive(true)
			So(err, ShouldBeNil)
		})
	})
}

func TestSetHvacMode(t *testing.T) {
	Convey("When setting the HvacMode", t, func() {
		client := New(ClientID, State, ClientSecret, AuthorizationCode)
		client.AccessTokenURL = ts.URL
		client.Authorize()
		client.Token = Token
		client.APIURL = ts.URL
		devices, _ := client.Devices()
		Convey("When an invalid mode given it should trow an error", func() {
			err := devices.Thermostats["z1234"].SetHvacMode(2000)
			So(err.Description, ShouldEqual, "Invalid HvacMode requested - must be cool, heat, heat-cool or off")
		})
		Convey("When requesting HvacMode off", func() {
			err := devices.Thermostats["z1234"].SetHvacMode(Off)
			So(err, ShouldBeNil)
		})
		Convey("When requesting HvacMode cool", func() {
			err := devices.Thermostats["z1234"].SetHvacMode(Cool)
			So(err, ShouldBeNil)
		})
		Convey("When requesting HvacMode heat", func() {
			err := devices.Thermostats["z1234"].SetHvacMode(Heat)
			So(err, ShouldBeNil)
		})
		Convey("When requesting HvacMode heat-cool", func() {
			err := devices.Thermostats["z1234"].SetHvacMode(HeatCool)
			So(err, ShouldBeNil)
		})
	})
}

func TestTargetTemps(t *testing.T) {
	Convey("When setting target temperatures", t, func() {
		client := New(ClientID, State, ClientSecret, AuthorizationCode)
		client.AccessTokenURL = ts.URL
		client.Authorize()
		client.Token = Token
		client.APIURL = ts.URL
		devices, _ := client.Devices()
		Convey("When requesting to set a target high low temperature", func() {
			Convey("When farenheit", func() {
				err := devices.Thermostats["z1234"].SetTargetTempHighLowF(75, 65)
				So(err, ShouldBeNil)
			})
			Convey("When celcius", func() {
				err := devices.Thermostats["z1234"].SetTargetTempHighLowC(25.2, 12.5)
				So(err, ShouldBeNil)
			})
		})
		Convey("When requesting in celcius", func() {
			err := devices.Thermostats["z1234"].SetTargetTempC(28.5)
			So(err, ShouldBeNil)
		})
		Convey("When requesting in farenheit", func() {
			err := devices.Thermostats["z1234"].SetTargetTempF(50)
			So(err, ShouldBeNil)
		})
		Convey("When an invalid target temperature in farenheit", func() {
			err := devices.Thermostats["z1234"].SetTargetTempF(600)
			So(err.Description, ShouldEqual, "Temperature must be between 50 and 90 Farenheit")
			err = devices.Thermostats["z1234"].SetTargetTempC(8)
			So(err.Description, ShouldEqual, "Temperature must be between 9 and 32 Celcius")
		})
	})
}
