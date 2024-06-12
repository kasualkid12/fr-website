import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './styles/App.scss';
import NavBar from './components/NavBar';
import Home from './components/Home';
import FamilyTree from './components/FamilyTree';

function App() {
  return (
    <div className="App">
      <Router>
        <NavBar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/family-tree" element={<FamilyTree />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
