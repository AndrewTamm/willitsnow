package response

import (
	"github.com/AndrewTamm/WillItSnow/cmd/model"
	"time"
)

// actual conditions: https://www.weatherapi.com/docs/weather_conditions.json
var iconMap = map[int]string{
	1000: "sunny",
	1003: "partly_cloudy",
	1006: "cloudy",
	1009: "overcast",
	1030: "mist",
	1063: "light_rain",
	1066: "light_snow",
	1069: "light_sleet",
	1072: "light_freezing_drizzle",
	1087: "thunder",
	1114: "blowing_snow",
	1117: "blizzard",
	1135: "fog",
	1147: "freezing_fog",
	1150: "light_drizzle",
	1153: "light_drizzle",
	1168: "freezing_drizzle",
	1171: "freezing_drizzle",
	1180: "light_rain",
	1183: "light_rain",
	1186: "rain",
	1189: "rain",
	1192: "heavy_rain",
	1195: "heavy_rain",
	1198: "light_freezing_rain",
	1201: "freezing_rain",
	1204: "light_sleet",
	1207: "sleet",
	1210: "patchy_light_snow",
	1213: "light_snow",
	1216: "snow",
	1219: "snow",
	1222: "heavy_snow",
	1225: "heavy_snow",
	1237: "ice_pellets",
	1240: "light_rain",
	1243: "rain",
	1246: "heavy_rain",
	1249: "light_sleet",
	1252: "sleet",
	1255: "light_snow",
	1258: "snow",
	1261: "light_ice_pellets",
	1264: "ice_pellets",
	1273: "thunder",
	1276: "thunder",
	1279: "thunder_snow",
	1283: "thunder_snow",
}

type Condition struct {
	Text string `json:"text"`
	Code int    `json:"code"`
}

type WeatherApiResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzId           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Forecast struct {
		DailyForecasts []struct {
			Date         string    `json:"date"`
			DateEpoch    int64     `json:"date_epoch"`
			Condition    Condition `json:"condition"`
			DailySummary struct {
				MaxTempF  float32   `json:"maxtemp_f"`
				MinTempF  float32   `json:"mintemp_f"`
				Condition Condition `json:"condition"`
			} `json:"day"`
			HourlyForecast []struct {
				Time         string    `json:"time"`
				TimeEpoch    int64     `json:"time_epoch"`
				Condition    Condition `json:"condition"`
				WillItRain   int       `json:"will_it_rain"`
				ChanceOfRain int       `json:"chance_of_rain"`
				WillItSnow   int       `json:"will_it_snow"`
				ChanceOfSnow int       `json:"chance_of_snow"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func (w WeatherApiResponse) AsModel() *model.Weather {
	var result model.Weather

	for _, day := range w.Forecast.DailyForecasts {
		dailyForecast := model.Forecast{
			TemperatureMax: day.DailySummary.MaxTempF,
			TemperatureMin: day.DailySummary.MinTempF,
			Time:           time.Unix(day.DateEpoch, 0),
			Icon:           weatherCondition(day.DailySummary.Condition),
		}

		for _, hour := range day.HourlyForecast {
			willItRain := hour.WillItRain == 1
			dailyForecast.WillItRain = dailyForecast.WillItRain || willItRain
			result.WillItRain = result.WillItRain || willItRain

			willItSnow := hour.WillItSnow == 1
			dailyForecast.WillItSnow = dailyForecast.WillItSnow || willItSnow
			result.WillItSnow = result.WillItSnow || willItSnow
		}
		result.WillItBeCold = result.WillItBeCold || day.DailySummary.MinTempF < 32
		result.Daily = append(result.Daily, dailyForecast)
	}

	return &result
}

func weatherCondition(condition Condition) string {
	if iconName, ok := iconMap[condition.Code]; ok {
		return iconName
	}
	return "unknown"
}
