import React from 'react';
import { render, screen } from '@testing-library/react';
import Footer from './Footer';

test('renders created by', () => {
  render(<Footer />);
  const linkElement = screen.getByText(/Created by/i);
  expect(linkElement).toBeInTheDocument();
});

test('renders inspiration link', () => {
  render(<Footer />);
  const inspiredElement = screen.getByText(/Inspired by/i);
  expect(inspiredElement).toBeInTheDocument();

  const linkElement = inspiredElement.querySelector('a');
  expect(linkElement).toHaveAttribute('href', 'https://www.willitsnowinpdx.com/');
});
