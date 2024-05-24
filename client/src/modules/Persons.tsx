import React from 'react';
import { Person } from '../interfaces/Person';

interface PersonsComponentProps {
  person: Person;
}

function PersonsComponent({ person }: PersonsComponentProps) {
  return (
    <div>
      <div key={person.id}>
        <h1>{person.name}</h1>
        <p>Birth Date: {person.birthDate}</p>
        <p>Death Date: {person.deathDate}</p>
        <p>Gender: {person.gender}</p>
        <p>Profile ID: {person.profileId}</p>
        <p>Relationship: {person.relationship}</p>
        {person.photoUrl && (
          <img src={person.photoUrl} alt={`${person.name}`} />
        )}
      </div>
    </div>
  );
}

export default PersonsComponent;
