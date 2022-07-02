import React, {useState} from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { MantineProvider } from '@mantine/core';
import {
    Register,
    Login,
    AppHeader,
    Home,
    StartEndTrip,
    City,
    ResumeTrip,
    MockTest,
    Travel,
    TripParent
} from "./components/organizationComponent";
import reportWebVitals from './reportWebVitals';

if (process.env.NODE_ENV === 'development') {
  const { worker } = require('./mocks/browser')
  worker.start()
}

function App() {

    let TXT = "Hello World";

    const [startCity, setStartCity] = useState('Paris');
    const [endCity, setEndCity] = useState('');
    const [startDateValue, setStartDate] = React.useState<Date | null>(new Date());
    const [endDateValue, setEndDate] = React.useState<Date | null>(new Date());

    let fulTrip = {
        startCity,
        setStartCity,
        endCity,
        setEndCity,
        startDateValue,
        setStartDate,
        endDateValue,
        setEndDate
    }


    return (
    <BrowserRouter>
      <AppHeader />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route path="/startEndTrip" element={<StartEndTrip fulTrip={fulTrip} />} />
        <Route path="/trip" element={<TripParent  />} />
        <Route path="/travel" element={<Travel txt={TXT}/>} />
        <Route path="/city" element={<City />} />
        <Route path="/resumeTrip" element={<ResumeTrip />} />
        <Route path="/mocking" element={<MockTest />} />

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
