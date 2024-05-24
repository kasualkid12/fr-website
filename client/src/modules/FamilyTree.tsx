import React, { useEffect, useState } from 'react';
import '../styles/FamilyTree.scss';
import PersonsComponent from './Persons';
import { Person } from '../interfaces/Person';

function FamilyTree() {
  const [persons, setPersons] = useState<Person[]>([]);

  useEffect(() => {
    fetch('http://localhost:8080/persons')
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then((data) => setPersons(data))
      .catch((error) => console.error('Error fetching data:', error));
  }, []);

  return (
    <div className="FamilyTree">
      {persons.map((person) => (
        <PersonsComponent key={person.id} person={person} />
      ))}
    </div>
  );
}

export default FamilyTree;
