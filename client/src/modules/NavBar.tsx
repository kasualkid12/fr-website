import React from 'react';

import '../styles/NavBar.css';
import crest from '../public/Hershey Crest.512x(Transparent).png';
import defaultPfp from '../public/Default user.svg'

function NavBar() {
  return (
    <div className="Navbar">
      <div className="FamilySeal">
        <img className="Crest" alt="The family crest" src={crest} />
      </div>
      <div className='Title'>Hershey Reunion</div>
      <a href="./Home.tsx">Home</a>
      <a href="./FamilyTree.tsx">Family Tree</a>
      <div className="ProfileImage">
        <img className='Pfp' alt='pfp' src={defaultPfp} />
      </div>
    </div>
  );
}

export default NavBar;
