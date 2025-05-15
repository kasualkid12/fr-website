import React from 'react';
import LandingPage from './LandingPage';
import AboutPage from './AboutPage';
import Schedule from './Schedule';
import Hymn from './Hymn';

function Home() {
  return <div className="Home">
    <LandingPage />
    <AboutPage />
    <Schedule />
    <Hymn />
  </div>;
}

export default Home;