package nest

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDevicesStream(t *testing.T) {
	Convey("When requesting a devices stream we should get two events", t, func() {
		client := New(ClientID, State, ClientSecret, AuthorizationCode)
		client.Authorize()
		client.Token = Token
		client.APIURL = ts.URL
		cnt := 0
		devicesChan := make(chan *Devices)
		go func() {
			client.DevicesStream(func(devices *Devices, err error) {
				cnt++
				devicesChan <- devices
			})
		}()
		for i := 0; i < 2; i++ {
			devices := <-devicesChan
			checkFields(devices)
			So(devices.Thermostats["z1234"].StructureID, ShouldEqual, "s1234")
			So(client.Token, ShouldEqual, devices.Thermostats["z1234"].Client.Token)
			So(client.Token, ShouldEqual, devices.SmokeCoAlarms["a1234"].Client.Token)
		}
		So(cnt, ShouldEqual, 2)
	})
}
