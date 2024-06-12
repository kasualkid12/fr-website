import React, { useEffect, useState, useRef } from 'react';
import '../styles/FamilyTree.scss';
import '../styles/FamilyTreeDesktop.scss';
import '../styles/FamilyTreeMobile.scss';
import PersonBubble from './PersonBubble';
import { Person, PersonWithSpouse } from '../interfaces/Person';

function FamilyTree() {
  const [persons, setPersons] = useState<Person[]>([]);
  const [selectedPersonId, setSelectedPersonId] = useState<number>(1);
  const [history, setHistory] = useState<number[]>([]);
  const svgRef = useRef<SVGSVGElement>(null);

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

  useEffect(() => {
    renderLines();
  }, [persons]);

  const handlePersonClick = (id: number) => {
    setHistory((prevHistory) => [...prevHistory, selectedPersonId]);
    setSelectedPersonId(id);
  };

  const handleGoBack = () => {
    setHistory((prevHistory) => {
      const newHistory = [...prevHistory];
      const previousId = newHistory.pop();
      if (previousId !== undefined) {
        setSelectedPersonId(previousId);
      }
      return newHistory;
    });
  };

  const handleGoToTop = () => {
    setHistory([]);
    setSelectedPersonId(1);
  };

  const createPersonBubblesDesktop = (persons: Person[]) => {
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
          <PersonBubble
            person={sourcePerson}
            spouse={sourcePerson.spouse}
            onClick={() => handlePersonClick(sourcePerson!.id)}
            isSelf={true}
            key={`${sourcePerson.firstName} ${sourcePerson.lastName}-${sourcePerson.id}`}
          />
        </div>
      );

      const childBubbles = children.map((child) => (
        <div className="child-bubble" key={child.id} id={`person-${child.id}`}>
          <PersonBubble
            person={child}
            spouse={child.spouse}
            onClick={() => handlePersonClick(child.id)}
            isSelf={false}
            key={`${child.firstName} ${child.lastName}-${child.id}`}
          />
        </div>
      ));

      bubbles.push(<div className="children">{childBubbles}</div>);
    }

    return bubbles;
  };

  const renderLines = () => {
    const svgElement = svgRef.current;
    if (!svgElement) return;

    const svgNS = 'http://www.w3.org/2000/svg'; // SVG namespace
    svgElement.innerHTML = ''; // Clear existing lines
    const parentElement = document.getElementById(`person-${selectedPersonId}`);
    if (!parentElement) return;

    const parentRect = parentElement.getBoundingClientRect();

    persons.forEach((person, index) => {
      if (index === 0) return;
      const child = persons[index];
      if (!child.relationship.includes('Child')) return;
      const childElement = document.getElementById(`person-${child.id}`);
      if (!childElement) return;
      const childRect = childElement.getBoundingClientRect();

      const line = document.createElementNS(svgNS, 'line');
      line.setAttribute(
        'x1',
        (parentRect.left + parentRect.width / 2).toString()
      );
      line.setAttribute(
        'y1',
        (parentRect.top + window.scrollY + parentRect.height / 2).toString()
      );
      line.setAttribute(
        'x2',
        (childRect.left + childRect.width / 2).toString()
      );
      line.setAttribute(
        'y2',
        (childRect.top + window.scrollY + childRect.height / 2).toString()
      );
      line.setAttribute('stroke', '#faf9f6');
      line.setAttribute('stroke-width', '2');
      svgElement.appendChild(line);
    });
  };

  return (
    <div className="FamilyTree">
      <button className="go-to-top" onClick={handleGoToTop}>
        Go to Top
      </button>
      {history.length > 0 && (
        <button className="go-back" onClick={handleGoBack}>
          ↑
        </button>
      )}
      {createPersonBubblesDesktop(persons)}
      {/* add mobile view */}
      <svg className="lines" ref={svgRef}>
        {/* Lines will be appended here */}
      </svg>
    </div>
  );
}

export default FamilyTree;
