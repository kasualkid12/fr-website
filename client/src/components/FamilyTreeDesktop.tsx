import React, { useEffect } from 'react';
import '../styles/FamilyTreeDesktop.scss';
import PersonBubble from './PersonBubble';
import { Person, ViewProps, PersonWithSpouse } from '../interfaces/Person';

/**
 * FamilyTreeDesktop component renders a family tree with person bubbles and lines connecting them.
 * It takes an array of persons, a function to handle person click, a reference to an SVG element where the lines will be rendered, and an id of the currently selected person.
 * It returns a div containing the person bubbles and the SVG element.
 *
 * @param {ViewProps} props - The props object containing the persons, handlePersonClick, svgRef, and selectedPersonId.
 * @returns {JSX.Element} - The FamilyTreeDesktop component.
 */
function FamilyTreeDesktop({
  persons, // The array of persons to render
  selectedPersonId, // The id of the currently selected person
  history, // The history of selected person ids
  handlePersonClick, // The function to handle person click
  handleGoBack, // The function to handle go back click
  handleGoToTop, // The function to handle go to top click
  svgRef, // The reference to the SVG element where the lines will be rendered
}: ViewProps): JSX.Element {
  // Render the lines whenever the persons array changes
  useEffect(() => {
    renderLines();
  }, [persons]);

  /**
   * Create person bubbles for the desktop view.
   * It takes an array of persons and returns an array of JSX elements representing the person bubbles.
   *
   * @param {Person[]} persons - The array of persons to create bubbles for.
   * @returns {JSX.Element[]} - The array of JSX elements representing the person bubbles.
   */
  const createPersonBubbles = (persons: Person[]): JSX.Element[] => {
    const bubbles: JSX.Element[] = [];
    let sourcePerson: PersonWithSpouse | null = null;
    const children: PersonWithSpouse[] = [];

    // Iterate through the persons array and create bubbles for each person
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

    // Create a bubble for the source person and add it to the bubbles array
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
            key={sourcePerson.id}
          />
        </div>
      );

      // Create bubbles for the children and add them to the bubbles array
      const childBubbles = children.map((child) => (
        <div className="child-bubble" key={child.id} id={`person-${child.id}`}>
          <PersonBubble
            person={child}
            spouse={child.spouse}
            onClick={() => handlePersonClick(child.id)}
            isSelf={false}
            key={child.id}
          />
        </div>
      ));

      bubbles.push(<div className="children">{childBubbles}</div>);
    }

    return bubbles;
  };

  /**
   * Render the lines connecting the person bubbles.
   * It calculates the coordinates of the lines and appends them to the SVG element.
   */
  const renderLines = (): void => {
    const svgElement = svgRef.current;
    if (!svgElement) return;

    const svgNS = 'http://www.w3.org/2000/svg'; // SVG namespace
    svgElement.innerHTML = ''; // Clear existing lines
    const parentElement = document.getElementById(`person-${selectedPersonId}`);
    if (!parentElement) return;

    const parentRect = parentElement.getBoundingClientRect();

    // Iterate through the persons array and create lines for each child
    persons.forEach((person, index) => {
      if (index === 0) return;
      const child = person;
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

  // Return the FamilyTreeDesktop component containing the person bubbles and the SVG element
  return (
    <div className="FamilyTreeDesktop">
      {/* Go to Top button */}
      <button className="go-to-top" onClick={handleGoToTop}>
        Top of Tree
      </button>
      {/* Go Back button */}
      {history.length > 0 && (
        <button className="go-back" onClick={handleGoBack}>
          ↑
        </button>
      )}
      {createPersonBubbles(persons)}
      <svg className="lines" ref={svgRef}>
        {/* Lines will be appended here */}
      </svg>
    </div>
  );
}

export default FamilyTreeDesktop;
