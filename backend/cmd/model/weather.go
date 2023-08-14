package model

import "time"

type Forecast struct {
	Icon           string    `json:"icon"`
	Time           time.Time `json:"time"`
	TemperatureMin float32   `json:"temperatureMin"`
	TemperatureMax float32   `json:"temperatureMax"`
	WillItSnow     bool      `json:"willItSnow"`
	WillItRain     bool      `json:"willItRain"`
}

type Weather struct {
	WillItSnow   bool       `json:"willItSnow"`
	WillItRain   bool       `json:"willItRain"`
	WillItBeCold bool       `json:"willItBeCold"`
	Daily        []Forecast `json:"daily"`
}
