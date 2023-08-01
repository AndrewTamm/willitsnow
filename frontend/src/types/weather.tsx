export interface DailyWeather {
    icon: string;
    time: string;
    temperatureMin: number;
    temperatureMax: number;
    willItSnow: boolean;
    willItRain: boolean;
}


export interface WeatherData {
    willItSnow: boolean;
    willItRain: boolean;
    willItBeCold: boolean;
    daily: DailyWeather[];
  }

