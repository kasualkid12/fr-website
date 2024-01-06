import React from 'react';
import './styles/App.scss';
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
