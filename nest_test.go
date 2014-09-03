package nest

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"
)

const (
	ClientID             = "1234"
	State                = "STATE"
	ClientSecret         = "5678"
	BadAuthorizationCode = "ABCD1234"
	AuthorizationCode    = "EFGH5678"
	Token                = "c.9i8pWrv..."
)

var ts *httptest.Server

func TestNew(t *testing.T) {
	ts = serveHTTP(t)

	Convey("Given a client ID and state we should be able to create a new client", t, func() {
		client := New(ClientID, State, ClientSecret, BadAuthorizationCode)
		client.AccessTokenURL = ts.URL
		err := client.Authorize()
		So(client.ID, ShouldEqual, ClientID)
		So(client.State, ShouldEqual, State)
		So(client.Secret, ShouldEqual, ClientSecret)
		So(client.AuthorizationCode, ShouldEqual, BadAuthorizationCode)
		Convey("Given we gave the oauth2 API a bad authorization code we should get an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error, ShouldEqual, "oauth2_error")
			So(err.Description, ShouldEqual, "authorization code not found")
		})
	})
	Convey("Given a client ID and state we should be able to create a new client", t, func() {
		client := New(ClientID, State, ClientSecret, AuthorizationCode)
		client.AccessTokenURL = ts.URL
		err := client.Authorize()
		So(client.ID, ShouldEqual, ClientID)
		So(client.State, ShouldEqual, State)
		So(client.Secret, ShouldEqual, ClientSecret)
		So(client.AuthorizationCode, ShouldEqual, AuthorizationCode)
		Convey("Given we gave the oauth2 API a valid authorization code we get an access token back", func() {
			So(err, ShouldBeNil)
			So(client.Token, ShouldEqual, Token)
			So(client.ExpiresIn, ShouldEqual, 315360000)
		})
	})
}

func TestDevices(t *testing.T) {
	Convey("When requesting a devices listing we should get a valid set of devices", t, func() {
		client := New(ClientID, State, ClientSecret, AuthorizationCode)
		client.Authorize()
		client.Token = Token
		client.APIURL = ts.URL
		devices, err := client.Devices()
		So(err, ShouldBeNil)
		checkFields(devices)
		So(client.Token, ShouldEqual, devices.Thermostats["z1234"].Client.Token)
		So(client.Token, ShouldEqual, devices.SmokeCoAlarms["z5678"].Client.Token)
	})
}
