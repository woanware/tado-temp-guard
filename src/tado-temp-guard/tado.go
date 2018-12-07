package main

import "time"

const JSON_OVERLAY string = `{
	"setting": {
	  "type": "HEATING",
	  "power": "ON",
	  "temperature": {
		"celsius": %2.1f
	  }
	},
	"termination": {
		"type": "TADO_MODE"
	}
 }`

//
type Home struct {
	//Email string `json:"email"`
	Homes []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"homes"`
	// ID            string `json:"id"`
	// Locale        string `json:"locale"`
	// MobileDevices []struct {
	// 	DeviceMetadata struct {
	// 		Locale    string `json:"locale"`
	// 		Model     string `json:"model"`
	// 		OsVersion string `json:"osVersion"`
	// 		Platform  string `json:"platform"`
	// 	} `json:"deviceMetadata"`
	// 	ID       int `json:"id"`
	// 	Location struct {
	// 		AtHome          bool `json:"atHome"`
	// 		BearingFromHome struct {
	// 			Degrees int     `json:"degrees"`
	// 			Radians float64 `json:"radians"`
	// 		} `json:"bearingFromHome"`
	// 		RelativeDistanceFromHomeFence float64 `json:"relativeDistanceFromHomeFence"`
	// 		Stale                         bool    `json:"stale"`
	// 	} `json:"location"`
	// 	Name     string `json:"name"`
	// 	Settings struct {
	// 		GeoTrackingEnabled bool `json:"geoTrackingEnabled"`
	// 	} `json:"settings"`
	// } `json:"mobileDevices"`
	// Name     string `json:"name"`
	// Username string `json:"username"`
}

//
type Zones []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	// DateCreated time.Time `json:"dateCreated"`
	// DeviceTypes []string  `json:"deviceTypes"`
	// Devices     []struct {
	// 	DeviceType       string `json:"deviceType"`
	// 	SerialNo         string `json:"serialNo"`
	// 	ShortSerialNo    string `json:"shortSerialNo"`
	// 	CurrentFwVersion string `json:"currentFwVersion"`
	// 	ConnectionState  struct {
	// 		Value     bool      `json:"value"`
	// 		Timestamp time.Time `json:"timestamp"`
	// 	} `json:"connectionState"`
	// 	Characteristics struct {
	// 		Capabilities []string `json:"capabilities"`
	// 	} `json:"characteristics"`
	// 	BatteryState string   `json:"batteryState"`
	// 	Duties       []string `json:"duties"`
	// } `json:"devices"`
	// ReportAvailable bool `json:"reportAvailable"`
	// SupportsDazzle  bool `json:"supportsDazzle"`
	// DazzleEnabled   bool `json:"dazzleEnabled"`
	// DazzleMode      struct {
	// 	Supported bool `json:"supported"`
	// 	Enabled   bool `json:"enabled"`
	// } `json:"dazzleMode"`
	// OpenWindowDetection struct {
	// 	Supported        bool `json:"supported"`
	// 	Enabled          bool `json:"enabled"`
	// 	TimeoutInSeconds int  `json:"timeoutInSeconds"`
	// } `json:"openWindowDetection"`
}

type Zone struct {
	// TadoMode                       string      `json:"tadoMode"`
	// GeolocationOverride            bool        `json:"geolocationOverride"`
	// GeolocationOverrideDisableTime interface{} `json:"geolocationOverrideDisableTime"`
	// Preparation                    interface{} `json:"preparation"`
	Setting struct {
		Type        string `json:"type"`
		Power       string `json:"power"`
		Temperature struct {
			Celsius    float64 `json:"celsius"`
			Fahrenheit float64 `json:"fahrenheit"`
		} `json:"temperature"`
	} `json:"setting"`
	// OverlayType        interface{} `json:"overlayType"`
	// Overlay            interface{} `json:"overlay"`
	// OpenWindow         interface{} `json:"openWindow"`
	// NextScheduleChange struct {
	// 	Start   time.Time `json:"start"`
	// 	Setting struct {
	// 		Type        string `json:"type"`
	// 		Power       string `json:"power"`
	// 		Temperature struct {
	// 			Celsius    int `json:"celsius"`
	// 			Fahrenheit int `json:"fahrenheit"`
	// 		} `json:"temperature"`
	// 	} `json:"setting"`
	// } `json:"nextScheduleChange"`
	// Link struct {
	// 	State string `json:"state"`
	// } `json:"link"`
	// ActivityDataPoints struct {
	// 	HeatingPower struct {
	// 		Type       string    `json:"type"`
	// 		Percentage int       `json:"percentage"`
	// 		Timestamp  time.Time `json:"timestamp"`
	// 	} `json:"heatingPower"`
	// } `json:"activityDataPoints"`
	// SensorDataPoints struct {
	// 	InsideTemperature struct {
	// 		Celsius    float64   `json:"celsius"`
	// 		Fahrenheit float64   `json:"fahrenheit"`
	// 		Timestamp  time.Time `json:"timestamp"`
	// 		Type       string    `json:"type"`
	// 		Precision  struct {
	// 			Celsius    float64 `json:"celsius"`
	// 			Fahrenheit float64 `json:"fahrenheit"`
	// 		} `json:"precision"`
	// 	} `json:"insideTemperature"`
	// 	Humidity struct {
	// 		Type       string    `json:"type"`
	// 		Percentage int       `json:"percentage"`
	// 		Timestamp  time.Time `json:"timestamp"`
	// 	} `json:"humidity"`
	// } `json:"sensorDataPoints"`
}

//
type Overlay struct {
	Type    string `json:"type"`
	Setting struct {
		Type        string `json:"type"`
		Power       string `json:"power"`
		Temperature struct {
			Celsius    float64 `json:"celsius"`
			Fahrenheit float64 `json:"fahrenheit"`
		} `json:"temperature"`
	} `json:"setting"`
	Termination struct {
		Type                   string    `json:"type"`
		DurationInSeconds      int       `json:"durationInSeconds"`
		Expiry                 time.Time `json:"expiry"`
		RemainingTimeInSeconds int       `json:"remainingTimeInSeconds"`
		ProjectedExpiry        time.Time `json:"projectedExpiry"`
	} `json:"termination"`
}
