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

func TestStructuresStream(t *testing.T) {
	Convey("When requesting a structures stream we should get two events", t, func() {
		client := New(ClientID, State, ClientSecret, AuthorizationCode)
		client.Authorize()
		client.Token = Token
		client.APIURL = ts.URL
		cnt := 0
		structuresChan := make(chan map[string]*Structure)
		go func() {
			client.StructuresStream(func(structures map[string]*Structure, err error) {
				cnt++
				structuresChan <- structures
			})
		}()
		for i := 0; i < 2; i++ {
			structures := <-structuresChan
			So(structures["s1234"].StructureID, ShouldEqual, "s1234")
			So(structures["s1234"].SmokeCoAlarms[0], ShouldEqual, "a1234")
			So(structures["s1234"].Name, ShouldEqual, "Miramar")
		}
		So(cnt, ShouldEqual, 2)
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

func structuresEventJSON() []byte {
	return []byte(`{"path":"/structures","data":{"s1234":{"smoke_co_alarms":["a1234"],"name":"Miramar","country_code":"US","away":"away","thermostats":["t1234"],"structure_id":"s1234"},"s5678":{"smoke_co_alarms":["a5678","_0suCE5N0Gt9ARz4UZTtVCqpa4SSjzbd"],"name":"Miramar","country_code":"US","away":"away","thermostats":["t5678"],"structure_id":"s5678"}}}`)
}
