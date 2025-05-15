import React, { JSX, useState } from 'react';
import '../styles/FamilyTreeDesktop.scss';
import PersonBubble from './PersonBubble';
import { Person, ViewProps, PersonWithSpouse } from '../interfaces/Person';
import ImageCropperModal from './ImageCropperModal';

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
  fetchImage, // The function to fetch image from the backend
  adminMode = false,
}: ViewProps): JSX.Element {
  // Modal state
  const [cropperOpen, setCropperOpen] = useState(false);
  const [cropperImage, setCropperImage] = useState<string | null>(null);
  const [uploadTarget, setUploadTarget] = useState<{
    person: Person;
    type: 'self' | 'spouse';
  } | null>(null);
  const fileInputRef = React.useRef<HTMLInputElement>(null);

  /**
   * Create person bubbles for the desktop view.
   * It takes an array of persons and returns an array of JSX elements representing the person bubbles.
   *
   * @param {Person[]} persons - The array of persons to create bubbles for.
   * @returns {JSX.Element[]} - The array of JSX elements representing the person bubbles.
   */
  const createPersonBubbles = (persons: Person[]): JSX.Element[] => {
    const bubbles: JSX.Element[] = [];
    let selfPerson: PersonWithSpouse | null = null;
    const children: PersonWithSpouse[] = [];

    // Iterate through the persons array and create bubbles for each person
    for (let i = 0; i < persons.length; i++) {
      if (!selfPerson) {
        selfPerson = persons[i];
        if (
          i < persons.length - 1 &&
          persons[i + 1].relationship.includes('Spouse')
        ) {
          selfPerson = { ...selfPerson, spouse: persons[i + 1] };
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
    if (selfPerson) {
      bubbles.push(
        <div
          className="self"
          key={selfPerson.id}
          id={`person-${selfPerson.id}`}
        >
          <PersonBubble
            person={selfPerson}
            spouse={selfPerson.spouse}
            isSelf={true}
            key={selfPerson.id}
            fetchImage={fetchImage}
            adminMode={adminMode}
            onUploadPhoto={() => handleUpload(selfPerson!.id, 'self')}
            onRemovePhoto={() => handleRemove(selfPerson!.id, 'self')}
          />
          {adminMode && (
            <div className="admin-controls">
              <button onClick={() => handleUpload(selfPerson.id, 'self')}>
                Upload Self Photo
              </button>
              <button onClick={() => handleRemove(selfPerson.id, 'self')}>
                Remove Self Photo
              </button>
              {selfPerson.spouse && (
                <>
                  <button
                    onClick={() =>
                      handleUpload(selfPerson.spouse!.id, 'spouse')
                    }
                  >
                    Upload Spouse Photo
                  </button>
                  <button
                    onClick={() =>
                      handleRemove(selfPerson.spouse!.id, 'spouse')
                    }
                  >
                    Remove Spouse Photo
                  </button>
                </>
              )}
            </div>
          )}
          <div className="self-details-box">
            <div className="self-details">
              <p>
                {selfPerson.firstName} {selfPerson.lastName}
              </p>
              {/* TODO: Add self details */}
              <p>TODO: Add self details</p>
            </div>
            <div className="spouse-details">
              <p>
                {selfPerson.spouse?.firstName} {selfPerson.spouse?.lastName}
              </p>
              {/* TODO: Add spouse details */}
              <p>TODO: Add spouse details</p>
            </div>
          </div>
        </div>
      );

      // Create bubbles for the children and add them to the bubbles array
      const childBubbles = children.map((child) => (
        <div key={child.id} id={`person-${child.id}`}>
          <PersonBubble
            person={child}
            spouse={child.spouse}
            onClick={() => handlePersonClick(child.id)}
            isSelf={false}
            key={child.id}
            fetchImage={fetchImage}
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
  // const renderLines = (): void => {
  //   const svgElement = svgRef.current;
  //   if (!svgElement) return;

  //   const svgNS = 'http://www.w3.org/2000/svg'; // SVG namespace
  //   svgElement.innerHTML = ''; // Clear existing lines
  //   const parentElement = document.getElementById(`person-${selectedPersonId}`);
  //   if (!parentElement) return;

  //   const parentRect = parentElement.getBoundingClientRect();

  //   // Iterate through the persons array and create lines for each child
  //   persons.forEach((person, index) => {
  //     if (index === 0) return;
  //     const child = person;
  //     if (!child.relationship.includes('Child')) return;
  //     const childElement = document.getElementById(`person-${child.id}`);
  //     if (!childElement) return;
  //     const childRect = childElement.getBoundingClientRect();

  //     const line = document.createElementNS(svgNS, 'line');
  //     line.setAttribute(
  //       'x1',
  //       (parentRect.left + parentRect.width / 2).toString()
  //     );
  //     line.setAttribute(
  //       'y1',
  //       (parentRect.top + window.scrollY + parentRect.height / 2).toString()
  //     );
  //     line.setAttribute(
  //       'x2',
  //       (childRect.left + childRect.width / 2).toString()
  //     );
  //     line.setAttribute(
  //       'y2',
  //       (childRect.top + window.scrollY + childRect.height / 2).toString()
  //     );
  //     line.setAttribute('stroke', '#faf9f6');
  //     line.setAttribute('stroke-width', '2');
  //     svgElement.appendChild(line);
  //   });
  // };

  const handleUpload = (personId: number, type: 'self' | 'spouse') => {
    // Find the person object
    const person = findPersonById(personId);
    if (!person) return;
    setUploadTarget({ person, type });
    // Trigger file input
    if (fileInputRef.current) fileInputRef.current.value = '';
    fileInputRef.current?.click();
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;
    const reader = new FileReader();
    reader.onload = (ev) => {
      setCropperImage(ev.target?.result as string);
      setCropperOpen(true);
    };
    reader.readAsDataURL(file);
  };

  const handleCropperClose = () => {
    setCropperOpen(false);
    setCropperImage(null);
    setUploadTarget(null);
  };

  // Helper to find person by id
  const findPersonById = (id: number): Person | undefined => {
    for (const p of persons) {
      if (p.id === id) return p;
      if (p.relationship.includes('Spouse') && p.profileId === id) return p;
    }
    return undefined;
  };

  const handleRemove = (personId: number, type: 'self' | 'spouse') => {
    alert(`Remove for ${type} photo of person ${personId}`);
  };

  // Compute objectName for upload if needed
  const objectName = uploadTarget
    ? `person_${uploadTarget.person.id}_${Date.now()}.jpg`
    : undefined;

  // Return the FamilyTreeDesktop component containing the person bubbles and the SVG element
  return (
    <div className="FamilyTreeDesktop">
      <input
        type="file"
        accept="image/*"
        style={{ display: 'none' }}
        ref={fileInputRef}
        onChange={handleFileChange}
      />
      <ImageCropperModal
        open={cropperOpen}
        imageSrc={cropperImage}
        onClose={handleCropperClose}
        apiEndpoint={
          uploadTarget ? 'http://localhost:8080/minio/addobject' : undefined
        }
        objectName={objectName}
        personId={uploadTarget?.person.id}
        formData={
          uploadTarget && objectName
            ? {
                bucketName: 'test-bucket',
                objectName: objectName,
                contentType: 'image/jpeg',
              }
            : {}
        }
        onSuccess={() => {
          handleCropperClose();
          // Optionally refresh UI or show a message
        }}
        onError={(err) => {
          // Optionally show error
        }}
      />
      {/* Go Back button */}
      {history.length > 0 && (
        <button className="go-back" onClick={handleGoBack}>
          ↑
        </button>
      )}
      {/* Go to Top button */}
      {history.length > 0 && (
        <button className="go-to-top" onClick={handleGoToTop}>
          Top of Tree
        </button>
      )}
      <div className="family-tree-container">
        {createPersonBubbles(persons)}
      </div>
      <svg className="lines" ref={svgRef}>
        {/* Lines will be appended here */}
      </svg>
    </div>
  );
}

export default FamilyTreeDesktop;
