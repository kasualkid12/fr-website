import React, { useEffect, useState } from 'react';

interface Person {
  id: number;
  name: string;
  birthDate: string;
  deathDate: string | null;
  gender: string;
  photoUrl: string | null;
  profileId: number | null;
}

function PersonsComponent() {
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
    <div>
      {persons.map((person) => (
        <div key={person.id}>
          <h1>{person.name}</h1>
          <p>Birth Date: {person.birthDate}</p>
          <p>Death Date: {person.deathDate}</p>
          <p>Gender: {person.gender}</p>
          <p>Profile ID: {person.profileId}</p>
          {person.photoUrl && (
            <img src={person.photoUrl} alt={`${person.name}`} />
          )}
        </div>
      ))}
    </div>
  );
}

export default PersonsComponent;
