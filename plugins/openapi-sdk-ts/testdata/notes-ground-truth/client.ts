import type { NewNote, Note } from "./types";

export class ApiError extends Error {
  statusCode: number;
  body: string;

  constructor(statusCode: number, body: string) {
    super(`API error ${statusCode}: ${body}`);
    this.name = "ApiError";
    this.statusCode = statusCode;
    this.body = body;
  }
}

export class Client {
  baseUrl: string;
  private fetchFn: typeof fetch;

  constructor(baseUrl: string, fetchFn: typeof fetch = fetch.bind(globalThis)) {
    this.baseUrl = baseUrl;
    this.fetchFn = fetchFn;
  }

  async listNotes(): Promise<Note[]> {
    let url = `${this.baseUrl}/notes`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
    return resp.json() as Promise<Note[]>;
  }

  async createNote(body: NewNote): Promise<Note> {
    let url = `${this.baseUrl}/notes`;
    const resp = await this.fetchFn(url, {
      method: "POST",
      headers: { "Content-Type": "application/json", "Accept": "application/json" },
      body: JSON.stringify(body),
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
    return resp.json() as Promise<Note>;
  }

  async getStats(): Promise<void> {
    let url = `${this.baseUrl}/notes/stats`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
  }

  async listSummaries(): Promise<void> {
    let url = `${this.baseUrl}/notes/summaries`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
  }

  async listTags(): Promise<void> {
    let url = `${this.baseUrl}/notes/tags`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
  }

  async getNote(noteId: string): Promise<Note> {
    let url = `${this.baseUrl}/notes/${encodeURIComponent(String(noteId))}`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
    return resp.json() as Promise<Note>;
  }
}
