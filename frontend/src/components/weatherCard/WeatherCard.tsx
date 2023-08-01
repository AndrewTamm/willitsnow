import moment from 'moment';
import React from 'react';
import { DailyWeather } from '../../types/weather';
import { sprintf } from 'sprintf'
import './WeatherCard.css'

interface IProps {
    weather: DailyWeather
} 

class WeatherCard extends React.Component<IProps> {
    render(): React.ReactNode {
        return (
            <div className="wc col s12 m6 l4" data-testid="weather-card">
                <div className="card blue-grey darken-1 ">

                    { (this.props.weather.willItSnow || this.props.weather.willItRain) ?
                        <div className="card-image">
                            <img  src={this.props.weather.willItSnow ? "images/snow.gif" : this.props.weather.willItRain ? "images/rain.gif" : ""}/>
                            <span className="card-title outline-shadow">{this.formatDate(this.props.weather.time)}</span>
                        </div> :
                        <div className="card-image">
                            <img src={"images/nice.gif"} />
                            <span className="card-title outline-shadow">{this.formatDate(this.props.weather.time)}</span>
                        </div>
                    }
                    <div className="card-content white-text">
                        <div className="icon-wrap">
                            <i className={"wi " + this.weatherIcon(this.props.weather.icon)}></i>
                        </div>
                        <div className={"weather"}>
                            <div className={"temp low"}>
                                <div className={"label"}>Lo:</div>
                                <div className={"value"}>{this.props.weather.temperatureMin}</div>
                            </div>
                            <div className={"temp high"}>
                                <div className={"label"}>Hi:</div>
                                <div className={"value"}>{this.props.weather.temperatureMax}</div>
                            </div>
                        </div>

                    </div>
                </div>
            </div>
        )
    }

    formatDate(timestamp: string): string {
        const date = moment.utc(timestamp)
        return sprintf('%s %d', date.format('ddd'), date.date())
    }

    weatherIcon(iconName: string): string {
        const iconMap = new Map([
            ["blizzard", "wi-snow-wind"],
            ["blowing_snow", "wi-day-snow-wind"],
            ["fog", "wi-fog"],
            ["freezing_fog", "wi-fog"],
            ["freezing_rain", "wi-rain-mix"],
            ["heavy_rain", "wi-day-rain"],
            ["heavy_snow", "wi-snow-wind"],
            ["ice_pellets", "wi-day-hail"],
            ["light_drizzle", "wi-day-sprinkle"],
            ["light_freezing_drizzle", "wi-rain-mix"],
            ["light_freezing_rain", "wi-rain-mix"],
            ["light_ice_pellets", "wi-day-hail"],
            ["light_rain", "wi-day-showers"],
            ["light_sleet", "wi-day-sleet"],
            ["light_snow", "wi-wu-chancesnow"],
            ["mist", "wi-day-sprinkle"],
            ["overcast", "wi-day-sunny-overcast"],
            ["partly_cloudy", "wi-wu-partlycloudy"],
            ["patchy_light_snow", "wi-wu-chancesnow"],
            ["rain", "wi-day-rain"],
            ["sleet", "wi-day-sleet"],
            ["snow", "wi-day-snow"],
            ["sunny", "wi-day-sunny"],
            ["thunder_snow", "wi-day-snow-thunderstorm"],
            ["thunder", "wi-day-thunderstorm"],
            ["cloudy", "wi-cloudy"]
        ]);

        return iconMap.get(iconName) ?? "";
    }
}

export default WeatherCard