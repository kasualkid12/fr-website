import React from 'react';
import LandingPage from './LandingPage';
import AboutPage from './AboutPage';

function Home() {
  return <div className="Home">
    <LandingPage />
    <AboutPage />
  </div>;
}

export default Home;