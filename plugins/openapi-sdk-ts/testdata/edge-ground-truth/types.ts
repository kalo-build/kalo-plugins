export interface ClientModel {
  id?: string;
  owner?: Pet;
}

export interface ClientModelModel {
  wraps?: ClientModel;
}

export interface NewClient {
  name?: string;
}

export interface Pet {
  id: number;
  name: string;
  tag?: string;
}

export type Pets = Pet[];

export type Status = string;

export interface ApiOverview {
  "ssh-key-fingerprints"?: Record<string, unknown>;
  verifiable_password_authentication?: boolean;
  "x/y"?: string;
}
