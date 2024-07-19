import React from 'react';
import LandingPage from './LandingPage';
import AboutPage from './AboutPage';
import Hymn from './Hymn';

function Home() {
  return <div className="Home">
    <LandingPage />
    <AboutPage />
    <Hymn />
  </div>;
}

export default Home;