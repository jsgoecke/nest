package nest

import (
	"time"
)

// APIError represents an error object from the Nest API
type APIError struct {
	Error       string `json:"error,omitempty"`
	Description string `json:"error_description,omitempty"`
	Status      string
	StatusCode  int
}

// Client represents a client object
type Client struct {
	ID                string
	State             string
	AuthorizationCode string
	Secret            string
	Token             string
	ExpiresIn         int
	AccessTokenURL    string
	APIURL            string
	RedirectURL       string
}

// Access represents a Nest access token object
type Access struct {
	Token     string `json:"access_token,omitempty"`
	ExpiresIn int    `json:"expires_in,omitempty"`
}

/*
Combined represents an object for the entire API
including devices and structures
*/
type Combined struct {
	Devices    *Devices              `json:"devices,omitempty"`
	Structures map[string]*Structure `json:"structures,omitempty"`
}

// DevicesEvent represents an object returned by the REST streaming API
type DevicesEvent struct {
	Path string   `json:"path,omitempty"`
	Data *Devices `json:"data,omitempty"`
}

// StructuresEvent represents an object returned by the REST streaming API
type StructuresEvent struct {
	Path string                `json:"path,omitempty"`
	Data map[string]*Structure `json:"data,omitempty"`
}

// Devices represents devices that include thermostats and smokecoalarms
type Devices struct {
	Thermostats   map[string]*Thermostat   `json:"thermostats,omitempty"`
	SmokeCoAlarms map[string]*SmokeCoAlarm `json:"smoke_co_alarms,omitempty"`
}

/*
Thermostat represents a Nest thermostat object
https://developer.nest.com/documentation/api#thermostats
https://developer.nest.com/documentation/how-to-thermostats-object
*/
type Thermostat struct {
	DeviceID               string    `json:"device_id,omitempty"`
	Locale                 string    `json:"locale,omitempty"`
	SoftwareVersion        string    `json:"software_version,omitempty"`
	StructureID            string    `json:"structure_id,omitempty"`
	Name                   string    `json:"name,omitempty"`
	NameLong               string    `json:"name_long,omitempty"`
	LastConnection         time.Time `json:"last_connection,omitempty"`
	IsOnline               bool      `json:"is_online,omitempty"`
	CanCool                bool      `json:"can_cool,omitempty"`
	CanHeat                bool      `json:"can_heat,omitempty"`
	IsUsingEmergencyHeat   bool      `json:"is_using_emergency_heat,omitempty"`
	HasFan                 bool      `json:"has_fan,omitempty"`
	FanTimerActive         bool      `json:"fan_timer_active,omitempty"`
	FanTimerTimeout        time.Time `json:"fan_timer_timeout,omitempty"`
	HasLeaf                bool      `json:"has_leaf,omitempty"`
	TemperatureScale       string    `json:"temperature_scale,omitempty"`
	TargetTemperatureF     int       `json:"target_temperature_f,omitempty"`
	TargetTemperatureC     float32   `json:"target_temperature_c,omitempty"`
	TargetTemperatureHighF int       `json:"target_temperature_high_f,omitempty"`
	TargetTemperatureHighC float32   `json:"target_temperature_high_c,omitempty"`
	TargetTemperatureLowF  int       `json:"target_temperature_low_f,omitempty"`
	TargetTemperatureLowC  float32   `json:"target_temperature_low_c,omitempty"`
	AwayTemperatureHighF   int       `json:"away_temperature_high_f,omitempty"`
	AwayTemperatureHighC   float32   `json:"away_temperature_high_c,omitempty"`
	AwayTemperatureLowF    int       `json:"away_temperature_low_f,omitempty"`
	AwayTemperatureLowC    float32   `json:"away_temperature_low_c,omitempty"`
	HvacMode               string    `json:"hvac_mode,omitempty"`
	AmbientTemperatureF    int       `json:"ambient_temperature_f,omitempty"`
	AmbientTemperatureC    float32   `json:"ambient_temperature_c,omitempty"`
	Humidity               int       `json:"humidity,omitempty"`
	HvacState              string    `json:"hvac_state,omitempty"`
	WhereID                string    `json:"where_id,omitempty"`
	Client                 *Client
}

// Tempratures represents all of the possible temprature settings for a Nest thermostat
type Tempratures struct {
	TargetTemperatureF     int     `json:"target_temperature_f,omitempty"`
	TargetTemperatureC     float32 `json:"target_temperature_c,omitempty"`
	TargetTemperatureHighF int     `json:"target_temperature_high_f,omitempty"`
	TargetTemperatureHighC float32 `json:"target_temperature_high_c,omitempty"`
	TargetTemperatureLowF  int     `json:"target_temperature_low_f,omitempty"`
	TargetTemperatureLowC  float32 `json:"target_temperature_low_c,omitempty"`
}

/*
SmokeCoAlarm represents a Nest smokecoalarm object (Smoke & CO2 alarm)
https://developer.nest.com/documentation/api#smoke_co_alarms
https://developer.nest.com/documentation/how-to-smoke-co-alarms-object
*/
type SmokeCoAlarm struct {
	DeviceID        string    `json:"device_id,omitempty"`
	Locale          string    `json:"locale,omitempty"`
	SoftwareVersion string    `json:"software_version,omitempty"`
	StructureID     string    `json:"structure_id,omitempty"`
	Name            string    `json:"name,omitempty"`
	NameLong        string    `json:"name_long,omitempty"`
	LastConnection  time.Time `json:"last_connection,omitempty"`
	IsOnline        bool      `json:"is_online,omitempty"`
	BatteryHealth   string    `json:"battery_health,omitempty"`
	CoAlarmState    string    `json:"co_alarm_state,omitempty"`
	SmokeAlarmState string    `json:"smoke_alarm_state,omitempty"`
	UIColorState    string    `json:"ui_color_state,omitempty"`
	Client          *Client
}

/*
Structure represents a Next structure object
https://developer.nest.com/documentation/api#structures
https://developer.nest.com/documentation/how-to-structures-object
*/
type Structure struct {
	StructureID         string    `json:"structure_id,omitempty"`
	Thermostats         []string  `json:"thermostats,omitempty"`
	SmokeCoAlarms       []string  `json:"smoke_co_alarms,omitempty"`
	Away                string    `json:"away,omitempty"`
	Name                string    `json:"name,omitempty"`
	CountryCode         string    `json:"country_code,omitempty"`
	PeakPeriodStartTime time.Time `json:"peak_period_start_time,omitempty"`
	PeakPeriodEndTime   time.Time `json:"peak_period_end_time,omitempty"`
	TimeZone            string    `json:"time_zone,omitempty"`
	ETA                 *ETA      `json:"eta,omitempty"`
	Client              *Client
}

// Eta represents an eta object (estimated time of a arrival for a structure)
type ETA struct {
	TripID                      string    `json:"trip_id,omitempty"`
	EstimatedArrivalWindowBegin time.Time `json:"estimated_arrival_window_begin,omitempty"`
	EstimatedArrivalWindowEnd   time.Time `json:"estimated_arrival_window_end,omitempty"`
}
