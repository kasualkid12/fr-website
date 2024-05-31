import React from 'react';
import { Person } from '../interfaces/Person';

interface PersonsComponentProps {
  person: Person;
  spouse?: Person;
  onClick: () => void;
}

function PersonsComponent({ person, spouse, onClick }: PersonsComponentProps) {
  return (
    <div className="bubble" onClick={onClick} id={`person-${person.id}`}>
      <div className="bubble-content">
        {person.photoUrl && (
          <img src={person.photoUrl} alt={`${person.name}`} />
        )}
        {spouse && spouse.photoUrl && (
          <img src={spouse.photoUrl} alt={`${spouse.name}`} />
        )}
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
