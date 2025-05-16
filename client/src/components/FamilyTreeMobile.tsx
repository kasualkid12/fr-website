import React, { JSX, useState, useEffect } from 'react';
import '../styles/FamilyTreeMobile.scss';
import { Person, ViewProps, PersonWithSpouse } from '../interfaces/Person';
import defaultImage from '../public/Default Image.svg';
import PersonBox from './PersonBox';

function FamilyTreeMobile({
  persons,
  history,
  handlePersonClick,
  handleGoBack,
  handleGoToTop,
  fetchImage,
}: ViewProps): JSX.Element {
  // State for main person and spouse images
  const [mainPersonImage, setMainPersonImage] = useState<string | null>(null);
  const [spouseImage, setSpouseImage] = useState<string | null>(null);
  const [sourcePerson, setSourcePerson] = useState<PersonWithSpouse | null>(
    null
  );
  const [children, setChildren] = useState<PersonWithSpouse[]>([]);

  // Derive sourcePerson and children from persons
  useEffect(() => {
    let localSourcePerson: PersonWithSpouse | null = null;
    const localChildren: PersonWithSpouse[] = [];
    for (let i = 0; i < persons.length; i++) {
      if (!localSourcePerson) {
        localSourcePerson = persons[i];
        if (
          i < persons.length - 1 &&
          persons[i + 1].relationship.includes('Spouse')
        ) {
          localSourcePerson = { ...localSourcePerson, spouse: persons[i + 1] };
          i++;
        }
      } else if (persons[i].relationship.includes('Child')) {
        let child: PersonWithSpouse = persons[i];
        if (
          i < persons.length - 1 &&
          persons[i + 1].relationship.includes('Spouse')
        ) {
          child = { ...child, spouse: persons[i + 1] };
          i++;
        }
        localChildren.push(child);
      }
    }
    setSourcePerson(localSourcePerson);
    setChildren(localChildren);
  }, [persons]);

  // Fetch images for main person and spouse
  useEffect(() => {
    let isMounted = true;
    if (sourcePerson && sourcePerson.photoUrl) {
      fetchImage(sourcePerson.photoUrl, 'test-bucket').then((img) => {
        if (isMounted) setMainPersonImage(img);
      });
    } else {
      setMainPersonImage(null);
    }
    if (sourcePerson && sourcePerson.spouse && sourcePerson.spouse.photoUrl) {
      fetchImage(sourcePerson.spouse.photoUrl, 'test-bucket').then((img) => {
        if (isMounted) setSpouseImage(img);
      });
    } else {
      setSpouseImage(null);
    }
    return () => {
      isMounted = false;
    };
  }, [sourcePerson, fetchImage]);

  const childBoxes = children.map((child) => (
    <div className="child-box" key={child.id} id={`person-${child.id}`}>
      <PersonBox
        person={child}
        spouse={child.spouse}
        onClick={() => handlePersonClick(child.id)}
        fetchImage={fetchImage}
      />
    </div>
  ));

  return (
    <div className="FamilyTreeMobile">
      <button className="go-to-top" onClick={handleGoToTop}>
        Top of Tree
      </button>
      {sourcePerson && (
        <div
          className="main-box"
          key={sourcePerson.id}
          id={`person-${sourcePerson.id}`}
        >
          <div className="header-box">
            <p>
              {sourcePerson.firstName} {sourcePerson.lastName}
            </p>
            {history.length > 0 && (
              <button className="go-back" onClick={handleGoBack}>
                Go Back
              </button>
            )}
          </div>
          <div className="body-box">
            <div className="photo-box">
              <img
                className="person-image"
                src={mainPersonImage || defaultImage}
                alt={`${sourcePerson.firstName} ${sourcePerson.lastName}`}
              />
              {sourcePerson.spouse && (
                <div className="spouse-box">
                  <img
                    className="spouse-image"
                    src={spouseImage || defaultImage}
                    alt={`${sourcePerson.spouse.firstName} ${sourcePerson.spouse.lastName}`}
                  />
                  <div className="spouse-name">
                    <p>
                      {sourcePerson.spouse.firstName}{' '}
                      {sourcePerson.spouse.lastName}
                    </p>
                    <p>
                      {sourcePerson.spouse.birthDate}
                      {sourcePerson.spouse.deathDate
                        ? ` - ${sourcePerson.spouse.deathDate}`
                        : ''}
                    </p>
                  </div>
                </div>
              )}
            </div>
            <div className="details-box">
              <p className="details-header">Details</p>
              {/* TODO: Add person details */}
              <p className="details-text">TODO: Add person details</p>
            </div>
            <div className="children-box">
              <p>Children</p>
              {childBoxes}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default FamilyTreeMobile;
