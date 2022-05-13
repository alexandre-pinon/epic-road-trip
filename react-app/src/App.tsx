import React from 'react';
import logo from './logo.svg';
import './App.css';
import { MantineProvider } from '@mantine/core';

/*
export function label(name: string) {
	return `Hello ${name.toUppercase()}`;
}
*/

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}


/*
function App() {
  return (
    <MantineProvider
      theme={{
        // Override any other properties from default theme
        fontFamily: 'Open Sans, sans serif',
        spacing: { xs: 15, sm: 20, md: 25, lg: 30, xl: 40 },
      }}
    >
			<App />	
    </MantineProvider>
  );
}
*/

export default App;
