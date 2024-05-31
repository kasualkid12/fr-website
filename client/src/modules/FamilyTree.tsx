import React, { useEffect, useState } from 'react';
import '../styles/FamilyTree.scss';
import PersonsComponent from './Persons';
import { Person, PersonWithSpouse } from '../interfaces/Person';

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

  const createPersonBubbles = (persons: Person[]) => {
    const bubbles = [];
    let sourcePerson: PersonWithSpouse | null = null;
    const children: PersonWithSpouse[] = [];

    for (let i = 0; i < persons.length; i++) {
      if (!sourcePerson) {
        sourcePerson = persons[i];
        if (
          i < persons.length - 1 &&
          persons[i + 1].relationship.includes('Spouse')
        ) {
          sourcePerson = { ...sourcePerson, spouse: persons[i + 1] };
          i++; // Skip the next person since they are the spouse
        }
      } else if (persons[i].relationship.includes('Child')) {
        let child: PersonWithSpouse = persons[i];
        if (
          i < persons.length - 1 &&
          persons[i + 1].relationship.includes('Spouse')
        ) {
          child = { ...child, spouse: persons[i + 1] };
          i++; // Skip the next person since they are the spouse
        }
        children.push(child);
      }
    }

    if (sourcePerson) {
      bubbles.push(
        <div
          className="source-bubble"
          key={sourcePerson.id}
          id={`person-${sourcePerson.id}`}
        >
          <PersonsComponent
            person={sourcePerson}
            spouse={sourcePerson.spouse}
            onClick={() => handlePersonClick(sourcePerson!.id)}
          />
        </div>
      );

      const childBubbles = children.map((child) => (
        <div className="child-bubble" key={child.id} id={`person-${child.id}`}>
          <PersonsComponent
            person={child}
            spouse={child.spouse}
            onClick={() => handlePersonClick(child.id)}
          />
        </div>
      ));

      bubbles.push(<div className="children">{childBubbles}</div>);
    }

    return bubbles;
  };

  const renderLines = () => {
    const lines: JSX.Element[] = [];
    const parentElement = document.getElementById(`person-${selectedPersonId}`);
    if (!parentElement) return null;

    const parentRect = parentElement.getBoundingClientRect();

    persons.forEach((person, index) => {
      if (index === 0) return;
      const child = persons[index];
      if (!child.relationship.includes('Child')) return;
      const childElement = document.getElementById(`person-${child.id}`);
      if (!childElement) return;
      const childRect = childElement.getBoundingClientRect();

      lines.push(
        <line
          key={index}
          x1={parentRect.left + parentRect.width / 2}
          y1={parentRect.top + parentRect.height / 2}
          x2={childRect.left + childRect.width / 2}
          y2={childRect.top + childRect.height / 2}
          stroke="black"
          strokeWidth="2"
        />
      );
    });

    return lines;
  };

  return (
    <div className="FamilyTree">
      {createPersonBubbles(persons)}
      <svg className="lines">{renderLines()}</svg>
    </div>
  );
}

export default FamilyTree;
