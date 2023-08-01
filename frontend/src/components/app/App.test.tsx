import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import App from './App';
import { WeatherData } from '../../types/weather';

// Mock fetch
import fetchMock from 'jest-fetch-mock';
fetchMock.enableMocks();

const mockWeatherData: WeatherData = {
  willItSnow: true,
  willItRain: false,
  willItBeCold: false,
  daily: [
    {
      time: '2023-03-27T12:00:00Z',
      temperatureMin: 10,
      temperatureMax: 20,
      willItSnow: false,
      willItRain: false,
      icon: 'sunny'
    },
    {
      time: '2023-03-28T12:00:00Z',
      temperatureMin: 12,
      temperatureMax: 22,
      willItSnow: false,
      willItRain: false,
      icon: 'sunny'
    }
  ]
};

beforeEach(() => {
  fetchMock.resetMocks();
});

describe('App', () => {
  test('renders correctly and fetches data', async () => {
    fetchMock.mockResponseOnce(JSON.stringify(mockWeatherData));

    render(<App />);
    await waitFor(() => expect(fetchMock).toHaveBeenCalledTimes(1));
    await waitFor(() => expect(screen.getByText('Yes')).toBeInTheDocument());

    expect(screen.getByText('Yes')).toBeInTheDocument();
    expect(screen.getByText('see forecast')).toBeInTheDocument();
    expect(screen.getAllByTestId('weather-card').length).toBe(2);
  });
});
