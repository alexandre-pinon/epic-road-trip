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

});