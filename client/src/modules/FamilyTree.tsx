import React from 'react';

import '../styles/FamilyTree.scss';
import PersonsComponent from './Persons';

function FamilyTree() {
  return (
  <div className="FamilyTree">
    <PersonsComponent />
  </div>
  );
}

export default FamilyTree;
