export interface Person {
  id: number;
  name: string;
  birthDate: string;
  deathDate: string | null;
  gender: string;
  photoUrl: string | null;
  profileId: number | null;
  relationship: string;
}