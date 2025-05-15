import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './styles/App.scss';
import NavBar from './components/NavBar';
import Home from './components/Home';
import FamilyTree from './components/FamilyTree';
import History from './components/History';
import Login from './components/Login';

function App() {
  return (
    <div className="App">
      <Router>
        {/* <Login /> */}
        <NavBar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/family-tree" element={<FamilyTree />} />
          <Route path="/history" element={<History />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
