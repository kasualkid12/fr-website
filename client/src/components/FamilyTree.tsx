import React, { useEffect, useState, useRef } from 'react';
import '../styles/FamilyTree.scss';
import FamilyTreeDesktop from './FamilyTreeDesktop';
import FamilyTreeMobile from './FamilyTreeMobile';
import { Person } from '../interfaces/Person';

/**
 * FamilyTree component that renders the family tree.
 * It fetches the persons and handles the navigation between persons.
 */
function FamilyTree() {
  // State variables
  const [persons, setPersons] = useState<Person[]>([]); // Array of persons
  const [selectedPersonId, setSelectedPersonId] = useState<number>(1); // Id of the selected person
  const [history, setHistory] = useState<number[]>([]); // History of selected person ids
  const svgRef = useRef<SVGSVGElement>(null); // Reference to the SVG element
  const [isMobile, setIsMobile] = useState<boolean>(window.innerWidth <= 1023); // State to check if mobile

  /**
   * Fetches the persons from the backend.
   * @param {number} id - The id of the person to start fetching from.
   */
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

  // Fetch persons when selectedPersonId changes
  useEffect(() => {
    fetchPersons(selectedPersonId);
  }, [selectedPersonId]);

  /**
   * Adds a resize event listener to update the isMobile state variable
   * when the window is resized.
   */
  useEffect(() => {
    /**
     * Handles the resize event by updating the isMobile state variable
     * based on the current window width.
     */
    const handleResize = () => {
      setIsMobile(window.innerWidth <= 1023);
    };

    // Add the resize event listener
    window.addEventListener('resize', handleResize);

    // Clean up the resize event listener when the component unmounts
    return () => window.removeEventListener('resize', handleResize);
  }, []);

  /**
   * Handles the click event on a person.
   * Adds the selected person id to the history and updates the selected person id.
   * @param {number} id - The id of the person that was clicked.
   */
  const handlePersonClick = (id: number) => {
    setHistory((prevHistory) => [...prevHistory, selectedPersonId]);
    setSelectedPersonId(id);
  };

  /**
   * Handles the click event on the go back button.
   * Pops the last person id from the history and updates the selected person id.
   */
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

  /**
   * Handles the click event on the go to top button.
   * Resets the history and updates the selected person id to 1.
   */
  const handleGoToTop = () => {
    setHistory([]);
    setSelectedPersonId(1);
  };

  return (
    <div className="FamilyTree">
      {isMobile ? (
        <FamilyTreeMobile
          persons={persons}
          selectedPersonId={selectedPersonId}
          history={history}
          handlePersonClick={handlePersonClick}
          handleGoBack={handleGoBack}
          handleGoToTop={handleGoToTop}
          svgRef={svgRef}
        />
      ) : (
        <FamilyTreeDesktop
          persons={persons}
          selectedPersonId={selectedPersonId}
          history={history}
          handlePersonClick={handlePersonClick}
          handleGoBack={handleGoBack}
          handleGoToTop={handleGoToTop}
          svgRef={svgRef}
        />
      )}
    </div>
  );
}

export default FamilyTree;
