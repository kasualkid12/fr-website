import React, { useState, useEffect } from 'react';
import { PersonProps } from '../interfaces/Person';
import defaultImage from '../public/Default Image.svg';

function PersonBubble({
  person,
  spouse,
  onClick,
  isSelf,
  fetchImage,
}: PersonProps) {
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
            src={personImage || defaultImage}
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
