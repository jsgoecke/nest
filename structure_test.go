package nest

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStructures(t *testing.T) {
	Convey("When getting a map of structures", t, func() {
		client := New(ClientID, State, ClientSecret, AuthorizationCode)
		client.AccessTokenURL = ts.URL
		client.Authorize()
		client.Token = Token
		client.APIURL = ts.URL
		structures, err := client.Structures()
		So(err, ShouldBeNil)
		for _, value := range structures {
			So(value.Name, ShouldEqual, "Miramar")
		}
	})
}

func TestSetAway(t *testing.T) {
	Convey("When setting the Away status", t, func() {
		client := New(ClientID, State, ClientSecret, AuthorizationCode)
		client.AccessTokenURL = ts.URL
		client.Authorize()
		client.Token = Token
		client.APIURL = ts.URL
		structures, _ := client.Structures()
		Convey("When setting to away", func() {
			err := structures["h68sn..."].SetAway(Away)
			So(err, ShouldBeNil)
		})
		Convey("When setting an invalid away status", func() {
			err := structures["h68sn..."].SetAway(2000)
			So(err.Description, ShouldEqual, "Invalid Away requested - must be home, away or auto-away")
		})
	})
}

func structuresJSON() []byte {
	return []byte(`
		{
		    "h68sn...": {
		        "smoke_co_alarms": [
		            "R8CHkMLaJ_ge3_kCApUWMyqpa4SSjzbd"
		        ],
		        "name": "Miramar",
		        "country_code": "US",
		        "away": "home",
		        "thermostats": [
		            "1cf6CGEN..."
		        ],
		        "structure_id": "h68sn..."
		    },
		    "WeLo...": {
		        "smoke_co_alarms": [
		            "_0suCE5N0GsW...",
		            "_0suCE5N0Gt9..."
		        ],
		        "name": "Miramar",
		        "country_code": "US",
		        "away": "home",
		        "thermostats": [
		            "ZgPfn..."
		        ],
		        "structure_id": "WeLo..."
		    }
		}
		`)
}
