import React, { useState } from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { MantineProvider } from '@mantine/core';

import { Register, Login, AppHeader, Home, Travel, StartEndTrip, City, ResumeTrip, MockTest, Enjoy } from "./components/organizationComponent";
import reportWebVitals from './reportWebVitals';
import { Sleep } from './components/Sleep/Sleep';
import { Eat } from './components/Eat/Eat';
import { Drink } from './components/Drink/Drink';
import { EnjoyArrival } from './components/Enjoy/EnjoyArrival';
import { SleepArrival } from './components/Sleep/SleepArrival';
import { EatArrival } from './components/Eat/EatArrival';
import { DrinkArrival } from "./components/Drink/DrinkArrival"

// if (process.env.NODE_ENV === 'development') {
//   const { worker } = require('./mocks/browser')
//   worker.start()
// }

function App() {

  let TXT = "Hello World";

  const [startCity, setStartCity] = useState('');
  const [endCity, setEndCity] = useState('');
  const [startDateValue, setStartDate] = React.useState<Date | null>(new Date());
  const [endDateValue, setEndDate] = React.useState<Date | null>(new Date());
  const [selectedTravel, setSelectedTravel] = useState([]);
  const [enjoyName, setEnjoyName] = useState('');
  const [sleep, setSleep] = useState('');
  const [eat, setEat] = useState('');
  const [drink, setDrink] = useState('');

  // Arrival City
  const [enjoyArrival, setEnjoyArrival] = useState('');
  const [sleepArrival, setSleepArrival] = useState('');
  const [eatArrival, setEatArrival] = useState('');
  const [drinkArrival, setDrinkArrival] = useState('');

  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [userID, setUserID] = useState('');

  const auth = {
    isAuthenticated,
    setIsAuthenticated,
    userID,
    setUserID
  }

  let fulTrip = {
    startCity,
    setStartCity,
    endCity,
    setEndCity,
    startDateValue,
    setStartDate,
    endDateValue,
    setEndDate,
    selectedTravel,
    setSelectedTravel,
    enjoyName,
    setEnjoyName,
    sleep,
    setSleep,
    eat,
    setEat,
    drink,
    setDrink,
    enjoyArrival,
    setEnjoyArrival,
    sleepArrival,
    setSleepArrival,
    eatArrival, setEatArrival,
    drinkArrival, setDrinkArrival,
  }


  return (
    <BrowserRouter>
      <AppHeader auth={auth} />
      <Routes>
        <Route path="/" element={<Home fulTrip={fulTrip} auth={auth} />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login auth={auth} />} />
        <Route path="/startEndTrip" element={<StartEndTrip fulTrip={fulTrip} />} />
        <Route path="/travel" element={<Travel fulTrip={fulTrip} />} />
        <Route path="/city" element={<City />} />
        <Route path="/resumeTrip" element={<ResumeTrip fulTrip={fulTrip} auth={auth} />} />
        <Route path="/mocking" element={<MockTest />} />
        <Route path="/enjoy" element={<Enjoy fulTrip={fulTrip} />} />
        <Route path="/enjoyArrival" element={<EnjoyArrival fulTrip={fulTrip} />} />
        <Route path="/sleepArrival" element={<SleepArrival fulTrip={fulTrip} />} />
        <Route path="/eatArrival" element={<EatArrival fulTrip={fulTrip} />} />
        <Route path="/drinkArrival" element={<DrinkArrival fulTrip={fulTrip} />} />
        <Route path="/sleep" element={<Sleep fulTrip={fulTrip} />} />
        <Route path="/eat" element={<Eat fulTrip={fulTrip} />} />
        <Route path="/drink" element={<Drink fulTrip={fulTrip} />} />

        {/* <Route path="/login" element={<Login />} /> */}
      </Routes>
    </BrowserRouter>
  )
}

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    {/* <Listings title="EPI Listings" /> */}
    {/* <MantineProvider theme={{ colorScheme: 'dark' }} withGlobalStyles withNormalizeCSS> */}
    <App />
    {/* </MantineProvider> */}
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
