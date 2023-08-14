import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import WeatherCard from './WeatherCard';
import { DailyWeather } from '../../types/weather';

const mockWeather: DailyWeather = {
    time: '2023-03-27T12:00:00Z',
    temperatureMin: 10,
    temperatureMax: 20,
    willItSnow: false,
    willItRain: false,
    icon: 'sunny'
};

describe('WeatherCard', () => {
    test('renders correctly', () => {
        render(<WeatherCard weather={mockWeather} />);
        expect(screen.getByText('Mon 27')).toBeInTheDocument();
        expect(screen.getByText('Lo:')).toBeInTheDocument();
        expect(screen.getByText('Hi:')).toBeInTheDocument();
        expect(screen.getByText('10')).toBeInTheDocument();
        expect(screen.getByText('20')).toBeInTheDocument();
    });

    test('formats date correctly', () => {
        const weatherCard = new WeatherCard({ weather: mockWeather });
        const formattedDate = weatherCard.formatDate(mockWeather.time);
        expect(formattedDate).toBe('Mon 27');
    });

    test('maps weather icons correctly', () => {
        const weatherCard = new WeatherCard({ weather: mockWeather });
        const icon = weatherCard.weatherIcon(mockWeather.icon);
        expect(icon).toBe('wi-day-sunny');
    });
});
