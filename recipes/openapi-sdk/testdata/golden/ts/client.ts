import type { Pet, Pets } from "./types";

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

  async listPets(limit?: string): Promise<Pets> {
    let url = `${this.baseUrl}/pets`;
    const query = new URLSearchParams();
    if (limit !== undefined) query.set("limit", limit);
    if (query.toString()) url += `?${query.toString()}`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
    return resp.json() as Promise<Pets>;
  }

  async createPets(body: Pet): Promise<void> {
    let url = `${this.baseUrl}/pets`;
    const resp = await this.fetchFn(url, {
      method: "POST",
      headers: { "Content-Type": "application/json", "Accept": "application/json" },
      body: JSON.stringify(body),
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
  }

  async showPetById(petId: string): Promise<Pet> {
    let url = `${this.baseUrl}/pets/${encodeURIComponent(String(petId))}`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
    return resp.json() as Promise<Pet>;
  }
}
