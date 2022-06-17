import React from 'react'
import { render, screen } from '@testing-library/react'
import { City } from '../components/Trip/City';

test("checks city activities text", () => {
  render(<City />);

  const exploreCityName = screen.getByText(/Explorez "Ville de départ"/i);
  expect(exploreCityName).toBeInTheDocument();

  const suggestions = screen.getByText(/Suggestions/i);
  expect(suggestions).toBeInTheDocument();

  const seDivertir = screen.getByText(/Se divertir/i);
  expect(seDivertir).toBeInTheDocument();

  const seReposer = screen.getByText(/Se reposer/i);
  expect(seReposer).toBeInTheDocument();

  const seRestaurer = screen.getByText(/Se restaurer/i);
  expect(seRestaurer).toBeInTheDocument();

  const retour = screen.getByText(/Retour/i);
  expect(retour).toBeInTheDocument();

  const activityOne = screen.getByText(/Activity 1/i);
  expect(activityOne).toBeInTheDocument();

  const activityTwo = screen.getByText(/Activity 2/i);
  expect(activityTwo).toBeInTheDocument();

  const activityThree = screen.getByText(/Activity 3/i);
  expect(activityThree).toBeInTheDocument();

  const hotelOne = screen.getByText(/Hôtel 1/i);
  expect(hotelOne).toBeInTheDocument();

  const hotelTwo = screen.getByText(/Hôtel 2/i);
  expect(hotelTwo).toBeInTheDocument();

  const hotelThree = screen.getByText(/Hôtel 3/i);
  expect(hotelThree).toBeInTheDocument();

  const restaurantOne = screen.getByText(/Restaurant 1/i);
  expect(restaurantOne).toBeInTheDocument();

  const restaurantTwo = screen.getByText(/Restaurant 2/i);
  expect(restaurantTwo).toBeInTheDocument();

  const restaurantThree = screen.getByText(/Restaurant 3/i);
  expect(restaurantThree).toBeInTheDocument();

});