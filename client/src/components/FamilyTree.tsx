import React, {
  Suspense,
  use,
  useMemo,
  useEffect,
  useState,
  useRef,
} from 'react';
import '../styles/FamilyTree.scss';
import FamilyTreeDesktop from './FamilyTreeDesktop';
import FamilyTreeMobile from './FamilyTreeMobile';
import { Person } from '../interfaces/Person';
import defaultImage from '../public/Default Image.svg';

export async function fetchImage(
  objectName: string,
  bucketName: string
): Promise<string> {
  try {
    const response = await fetch('http://localhost:8080/minio/getobject', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        bucketName,
        objectName,
        contentType: 'image/jpeg',
      }),
    });
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    const blob = await response.blob();
    return URL.createObjectURL(blob);
  } catch (error) {
    console.error('Error fetching image for object:', objectName, error);
    // Fallback: you can either return a default URL or throw an error
    return defaultImage;
  }
}

/**
 * Async function to fetch persons from the backend.
 * Returns a Promise that resolves to an array of Person objects.
 */
async function fetchPersons(id: number): Promise<Person[]> {
  const response = await fetch(`http://localhost:8080/persons`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ id }),
  });
  if (!response.ok) {
    throw new Error('Network response was not ok');
  }
  return response.json();
}

// A simple cache to store promises by person id.
const personsCache = new Map<number, Promise<Person[]>>();

function getPersonsPromise(id: number): Promise<Person[]> {
  if (!personsCache.has(id)) {
    personsCache.set(id, fetchPersons(id));
  }
  return personsCache.get(id)!;
}

function usePersons(id: number): Person[] {
  // Memoize the promise so that if the id doesn't change, we reuse the promise.
  const personsPromise = useMemo(() => getPersonsPromise(id), [id]);
  return use(personsPromise);
}

/**
 * FamilyTree component that renders the family tree.
 * It fetches the persons and handles the navigation between persons.
 */
function FamilyTreeComponent() {
  // State variables
  const [selectedPersonId, setSelectedPersonId] = useState<number>(1);
  const [history, setHistory] = useState<number[]>([]);
  const svgRef = useRef<SVGSVGElement>(null);
  const [isMobile, setIsMobile] = useState<boolean>(window.innerWidth <= 1023);
  const [adminMode, setAdminMode] = useState<boolean>(false); // Admin mode toggle

  useEffect(() => {
    const handleResize = () => setIsMobile(window.innerWidth <= 1023);
    window.addEventListener('resize', handleResize);
    return () => window.removeEventListener('resize', handleResize);
  }, []);

  // Use the custom hook to get persons; this will suspend until the promise resolves.
  const persons = usePersons(selectedPersonId);

  /**
   * Handlers for navigating between persons.
   */
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

  return (
    <div className="FamilyTree">
      {/* Admin mode toggle (visible to all for now) */}
      <button
        style={{ position: 'absolute', top: 10, right: 10, zIndex: 100 }}
        onClick={() => setAdminMode((prev) => !prev)}
      >
        {adminMode ? 'Disable Admin Controls' : 'Enable Admin Controls'}
      </button>
      {isMobile ? (
        <FamilyTreeMobile
          persons={persons}
          selectedPersonId={selectedPersonId}
          history={history}
          handlePersonClick={handlePersonClick}
          handleGoBack={handleGoBack}
          handleGoToTop={handleGoToTop}
          svgRef={svgRef}
          fetchImage={fetchImage}
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
          fetchImage={fetchImage}
          adminMode={adminMode}
        />
      )}
    </div>
  );
}

/**
 * The FamilyTree component is wrapped in a Suspense boundary so that while the
 * async fetch (via the new `use` hook) is pending, a fallback UI is shown.
 */
export default function FamilyTree() {
  return (
    <Suspense fallback={<div>Loading family tree...</div>}>
      <FamilyTreeComponent />
    </Suspense>
  );
}
