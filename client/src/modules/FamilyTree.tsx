import React, { useEffect, useState } from 'react';
import '../styles/FamilyTree.scss';
import PersonsComponent from './Persons';
import { Person } from '../interfaces/Person';

function FamilyTree() {
  const [persons, setPersons] = useState<Person[]>([]);
  const [selectedPersonId, setSelectedPersonId] = useState<number>(1);

  const fetchPersons = (id: number) => {
    fetch(`http://localhost:8080/persons`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ id }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then((data) => setPersons(data))
      .catch((error) => console.error('Error fetching data:', error));
  };

  useEffect(() => {
    fetchPersons(selectedPersonId);
  }, [selectedPersonId]);

  const handlePersonClick = (id: number) => {
    setSelectedPersonId(id);
  };

  return (
    <div className="FamilyTree">
      {persons.map((person) => (
        <PersonsComponent
          key={person.id}
          person={person}
          onClick={() => handlePersonClick(person.id)}
        />
      ))}
    </div>
  );
}

export default FamilyTree;
