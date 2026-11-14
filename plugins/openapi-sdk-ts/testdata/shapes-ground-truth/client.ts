import type { ApiError, Item } from "./types";

export class ClientApiError extends Error {
  statusCode: number;
  body: string;

  constructor(statusCode: number, body: string) {
    super(`API error ${statusCode}: ${body}`);
    this.name = "ClientApiError";
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

  async listErrors(): Promise<ApiError> {
    let url = `${this.baseUrl}/errors`;
    const resp = await this.fetchFn(url, {
      method: "GET",
      headers: { "Accept": "application/json" },
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ClientApiError(resp.status, body);
    }
    return resp.json() as Promise<ApiError>;
  }

  async updateItem(id: string, body: Item, dryRun?: string): Promise<Item> {
    let url = `${this.baseUrl}/items/${encodeURIComponent(String(id))}`;
    const query = new URLSearchParams();
    if (dryRun !== undefined) query.set("dry_run", dryRun);
    if (query.toString()) url += `?${query.toString()}`;
    const resp = await this.fetchFn(url, {
      method: "PUT",
      headers: { "Content-Type": "application/json", "Accept": "application/json" },
      body: JSON.stringify(body),
    });
    if (!resp.ok) {
      const body = await resp.text();
      throw new ClientApiError(resp.status, body);
    }
    return resp.json() as Promise<Item>;
  }
}
