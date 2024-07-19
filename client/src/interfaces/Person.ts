import { RefObject } from 'react';

export interface Person {
  id: number;
  firstName: string;
  lastName: string;
  birthDate: string;
  deathDate: string | null;
  gender: string;
  photoUrl: string | null;
  profileId: number | null;
  relationship: string;
}

export interface PersonWithSpouse extends Person {
  spouse?: Person;
}

export interface PersonProps {
  person: Person;
  spouse?: Person;
  onClick?: () => void;
  isSelf?: boolean;
}

export interface ViewProps {
  persons: Person[];
  selectedPersonId: number;
  history: number[];
  handlePersonClick: (id: number) => void;
  handleGoBack: () => void;
  handleGoToTop: () => void;
  svgRef: RefObject<SVGSVGElement>;
}
