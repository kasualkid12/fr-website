import React, { useState } from 'react';
import '../styles/NavBar.scss';
import crest from '../public/Hershey Crest.512x(Transparent).png';
import defaultPfp from '../public/Default user.svg';

function NavBar() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="Navbar">
      <div className={`Hamburger ${isOpen ? 'open' : ''}`} onClick={() => setIsOpen(!isOpen)}>
        <div className="Bar"></div>
        <div className="Bar"></div>
        <div className="Bar"></div>
      </div>
      <div className="CenterContent">
        <div className="FamilySeal">
          <img className="Crest" alt="The family crest" src={crest} />
        </div>
        <div className="Title">Hershey Reunion</div>
      </div>
      <div className={`DropdownMenu ${isOpen ? 'open' : ''}`}>
          <a href="/">Home</a>
          <a href="/family-tree">Family Tree</a>
        </div>
      {/* Can be added back when accounts are added */}
      <div className="ProfileImage">
        <img className="Pfp" alt="pfp" src={defaultPfp} />
      </div>
    </div>
  );
}

export default NavBar;