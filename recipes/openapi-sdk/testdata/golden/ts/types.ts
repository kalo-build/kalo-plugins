export interface Error {
  code: number;
  message: string;
}

export interface Pet {
  id: number;
  name: string;
  tag?: string;
}

export type Pets = Pet[];
