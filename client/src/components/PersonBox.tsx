import React, { useState, useEffect } from 'react';
import { PersonProps } from '../interfaces/Person';
import defaultImage from '../public/Default Image.svg'; // Adjust the path as necessary

function PersonBox({ person, spouse, onClick, fetchImage }: PersonProps) {
  const [personImage, setPersonImage] = useState<string | null>(null);

  useEffect(() => {
    let isMounted = true;
    if (person.photoUrl) {
      fetchImage(person.photoUrl, 'test-bucket').then((image) => {
        if (isMounted) {
          setPersonImage(image);
        }
      });
    } else {
      setPersonImage(null);
    }
    return () => {
      isMounted = false;
    };
  }, [person.photoUrl, fetchImage]);

  if (!person) return null;

  return (
    <div className={`person-box`} onClick={onClick}>
      <img
        className="person-image"
        src={personImage || defaultImage}
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
