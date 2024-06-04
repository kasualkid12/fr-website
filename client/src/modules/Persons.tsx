import React from 'react';
import { Person } from '../interfaces/Person';
import defaultImage from '../public/Default Image.svg'; // Adjust the path as necessary

interface PersonsComponentProps {
  person: Person;
  spouse?: Person;
  onClick: () => void;
}

function PersonsComponent({ person, spouse, onClick }: PersonsComponentProps) {
  return (
    <div className="bubble" onClick={onClick} id={`person-${person.id}`}>
      <div className="bubble-content">
        <div className="images-container">
          <img src={person.photoUrl || defaultImage} alt={`${person.name}`} />
          {spouse && (
            <img src={spouse.photoUrl || defaultImage} alt={`${spouse.name}`} />
          )}
        </div>
        <div className="overlay">
          <p>
            {person.name} {spouse ? `& ${spouse.name}` : ''}
          </p>
          <p>
            {person.birthDate} - {person.deathDate}
          </p>
          {spouse && (
            <p>
              {spouse.birthDate} - {spouse.deathDate}
            </p>
          )}
        </div>
      </div>
    </div>
  );
}

export default PersonsComponent;
