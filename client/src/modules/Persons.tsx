import React from 'react';
import { Person } from '../interfaces/Person';
import defaultImage from '../public/Default Image.svg'; // Adjust the path as necessary

interface PersonsComponentProps {
  person: Person;
  spouse?: Person;
  onClick: () => void;
  isSelf?: boolean;
}

function PersonsComponent({
  person,
  spouse,
  onClick,
  isSelf,
}: PersonsComponentProps) {
  return (
    <div
      className={`bubble ${isSelf ? 'self-bubble' : 'child-bubble'}`}
      onClick={onClick}
      id={`person-${person.id}`}
    >
      <div className="bubble-content">
        <div className="images-container">
          <img
            className="person-image"
            src={person.photoUrl || defaultImage}
            alt={`${person.firstName} ${person.lastName}`}
          />
          {spouse && (
            <img
              className="spouse-image"
              src={spouse.photoUrl || defaultImage}
              alt={`${spouse.firstName} ${spouse.lastName}`}
            />
          )}
        </div>
        <div className="overlay">
          <p>
            {person.firstName} {spouse ? `& ${spouse.firstName}` : ''}{' '}
            {person.lastName}
          </p>
          <p>
            {isSelf
              ? `${person.birthDate} ${
                  person.deathDate ? `- ${person.deathDate}` : ''
                }`
              : ''}
          </p>
          <p>
            {isSelf && spouse
              ? `${spouse.birthDate} ${
                  spouse.deathDate ? `- ${spouse.deathDate}` : ''
                }`
              : ''}
          </p>
        </div>
      </div>
    </div>
  );
}

export default PersonsComponent;
