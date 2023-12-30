import React from 'react';

import '../styles/NavBar.css';

function NavBar() {
  return (
    <div className="Navbar">
      <div className="FamilySeal"></div>
      <a href='./Home.tsx'>Home</a>
      <a href='./FamilyTree.tsx'>Family Tree</a>
      <div className="profileImage"></div>
    </div>
  );
}

export default NavBar;
