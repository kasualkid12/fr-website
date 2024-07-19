import React from 'react';
import { PersonProps } from '../interfaces/Person';
import defaultImage from '../public/Default Image.svg'; // Adjust the path as necessary

function PersonBubble({ person, spouse, onClick, isSelf }: PersonProps) {
  return (
    <div
      className={`bubble ${isSelf ? 'self-bubble' : 'child-bubble'}`}
      onClick={isSelf ? undefined : onClick}
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
        {!isSelf ? (
          <div className="overlay">
            <p>
              {person.firstName} {spouse ? `& ${spouse.firstName}` : ''}{' '}
              {person.lastName}
            </p>
          </div>
        ) : null}
      </div>
    </div>
  );
}

export default PersonBubble;
