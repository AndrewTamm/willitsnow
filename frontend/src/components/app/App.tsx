import React from 'react';
import './App.css';
import { env } from '../../env/env'
import WeatherCard from '../weatherCard/WeatherCard'
import {WeatherData, DailyWeather} from '../../types/weather';


interface IProps {
}


interface IState {
  weather: WeatherData | null;
}

class App extends React.Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      weather: null
    } as IState;
  }

  componentDidMount(): void {


    fetch(env.weatherUrl)
        .then(response => response.json())
        .then(data => {
          this.setState({ weather: data });
        });
  }

  render(): JSX.Element {
    const {weather} = this.state

    return (
      <div className="App">
        <header className={`App-header ${weather?.willItSnow ? "red" : "light-green accent-4"}`}>
           <div className={"will-it-snow"}>
               { weather?.willItSnow ? "Yes" : "No" }
           </div>

           <div className={"bottom-container"}>
               <div className={"down-arrow bounce-on-load"} ></div>
             <div className={"cta"}>see forecast</div>
           </div>
        </header>
        <div className="row">
          {
            weather?.daily.map((dailyWeather: DailyWeather) => (
              <div className={"forecasts"} key={dailyWeather.time}>
                <WeatherCard weather={dailyWeather}/>
              </div>
            ))
          }
        </div>
      </div>
      
    )
  }
}

export default App;
