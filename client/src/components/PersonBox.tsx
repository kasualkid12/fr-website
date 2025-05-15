import React from 'react';
import { PersonProps } from '../interfaces/Person';
import defaultImage from '../public/Default Image.svg'; // Adjust the path as necessary

function PersonBox({ person, spouse, onClick, fetchImage }: PersonProps) {
  if (!person) return null;

  return (
    <div className={`person-box`} onClick={onClick}>
      <img
        className="person-image"
        src={person.photoUrl || defaultImage}
        alt={`${person.firstName} ${person.lastName}`}
      />
      <p className="person-name">
        {person.firstName} {spouse ? `& ${spouse.firstName}` : ''}{' '}
        {person.lastName}
      </p>
    </div>
  );
}

export default PersonBox;
