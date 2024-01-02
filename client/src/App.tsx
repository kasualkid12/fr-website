import React from 'react';
import logo from './public/logo.svg';
import './styles/App.css';
import NavBar from './modules/NavBar';
import Home from './modules/Home';

function App() {
  return (
    <div className="App">
      <header>
      <NavBar />
      </header>
      <Home />
    </div>
  );
}

export default App;
