import type { ApiOverview, Pets, Status } from "./types";

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

  async metaRoot(): Promise<ApiOverview> {
    let url = `${this.baseUrl}/meta/root`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
    return resp.json() as Promise<ApiOverview>;
  }

  async actionsCacheList(orgName: string, _2fa?: string): Promise<Status> {
    let url = `${this.baseUrl}/orgs/${encodeURIComponent(String(orgName))}/cache-usage`;
    const query = new URLSearchParams();
    if (_2fa !== undefined) query.set("2fa", _2fa);
    if (query.toString()) url += `?${query.toString()}`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
    return resp.json() as Promise<Status>;
  }

  async listPets(limit?: string, new_?: string): Promise<Pets> {
    let url = `${this.baseUrl}/pets`;
    const query = new URLSearchParams();
    if (limit !== undefined) query.set("limit", limit);
    if (new_ !== undefined) query.set("new", new_);
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

  async getStatus(): Promise<Status> {
    let url = `${this.baseUrl}/status`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
    return resp.json() as Promise<Status>;
  }

  async createStatus(body: ApiOverview): Promise<Status> {
    let url = `${this.baseUrl}/status`;
    const resp = await this.fetchFn(url, {
      method: "POST",
      headers: { "Content-Type": "application/json", "Accept": "application/json" },
      body: JSON.stringify(body),
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ApiError(resp.status, body);
    }
    return resp.json() as Promise<Status>;
  }
}
