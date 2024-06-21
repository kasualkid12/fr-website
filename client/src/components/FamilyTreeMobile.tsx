import React from 'react';
import '../styles/FamilyTreeMobile.scss';
import { Person, ViewProps, PersonWithSpouse } from '../interfaces/Person';
import defaultImage from '../public/Default Image.svg';

function FamilyTreeMobile({
  persons,
  history,
  handlePersonClick,
  handleGoBack,
  handleGoToTop,
}: ViewProps): JSX.Element {
  const createPersonBox = (persons: Person[]): JSX.Element => {
    let box: JSX.Element = <></>;
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
      const childBoxes = children.map((child) => (
        <div className="child-box" key={child.id} id={`person-${child.id}`}>
          <div
            className={`person-box`}
            onClick={() => handlePersonClick(child.id)}
          >
            <img
              className="person-image"
              src={child.photoUrl || defaultImage}
              alt={`${child.firstName} ${child.lastName}`}
            />
            <p className="person-name">
              {child.firstName}{' '}
              {child.spouse ? `& ${child.spouse.firstName}` : ''}{' '}
              {child.lastName}
            </p>
          </div>
        </div>
      ));

      box = (
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
                src={sourcePerson.photoUrl || defaultImage}
                alt={`${sourcePerson.firstName} ${sourcePerson.lastName}`}
              />
              {sourcePerson.spouse && (
                <div className="spouse-box">
                  <img
                    className="spouse-image"
                    src={sourcePerson.spouse.photoUrl || defaultImage}
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
      );
    }

    return box;
  };

  return (
    <div className="FamilyTreeMobile">
      <button className="go-to-top" onClick={handleGoToTop}>
        Top of Tree
      </button>
      {createPersonBox(persons)}
    </div>
  );
}

export default FamilyTreeMobile;
