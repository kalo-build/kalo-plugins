export interface NewNote {
  body?: string;
  title: string;
}

export interface Note {
  body?: string;
  createdAt?: string;
  id: string;
  pinned?: boolean;
  title: string;
}
