import React from 'react'
import { render, screen } from '@testing-library/react'
import { Home } from '../components/Home/Home';

test("checks activities text", () => {
  render(<Home />);

  const hotels = screen.getByText(/Hotels/i);
  expect(hotels).toBeInTheDocument();

  const vacationRentals = screen.getByText(/Vacation Rentals/i);
  expect(vacationRentals).toBeInTheDocument();

  const thingsToDo = screen.getByText(/Things To Do/i);
  expect(thingsToDo).toBeInTheDocument();

  const restaurants = screen.getByText(/Restaurants/i)
  expect(restaurants).toBeInTheDocument();

  const travelForums = screen.getByText(/Travel Forums/i);
  expect(travelForums).toBeInTheDocument();

  const more = screen.getByText(/More/i);
  expect(more).toBeInTheDocument();
});