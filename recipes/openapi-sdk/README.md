# recipe: openapi-sdk

OpenAPI 3.x file in, a TypeScript client and a Go client out. One command, no `kalo.yaml`.

Try it on kalo.build: [kalo.build/openapi-sdk](https://kalo.build/openapi-sdk).

> Technical preview. Built by an AI agent acting for Kalo Software GmbH.

```sh
kalo spec run openapi-sdk --input ./openapi.yaml --out ./sdk
# ./sdk/ts/{types,client,index}.ts   ./sdk/go/{types,client}.go
```

## Install (pre-release CLI, tested on Linux)

```sh
curl -fsSL https://raw.githubusercontent.com/kalo-build/kalo-cli/agent-main/install.sh | KALO_VERSION=v0.2.0-agent.4 KALO_INSTALL_DIR="$HOME/.local/bin" sh
export PATH="$HOME/.local/bin:$PATH"
kalo spec run openapi-sdk --input ./openapi.yaml --out ./sdk
```

No root needed. The first run compiles the two plugins and says so: allow about a minute on a small file. After that a run takes a few seconds.

### From a checkout (contributors, needs Go 1.24+)

```sh
scripts/stage-recipe.sh openapi-sdk            # builds both plugins to WASM, fills in digests
kalo spec run --recipe-dir dist/recipes/openapi-sdk --input ./openapi.yaml --out ./sdk
```

## What you get

- `ts/`: `interface` per object schema, `type` alias for every other component schema (arrays, enums, scalars, unions), a `fetch`-based `Client` class with one method per operation.
- `go/`: struct per object schema, defined type for every other component schema, a `Client` with one method per operation (`context.Context` first).
- Nothing else is written to `--out`. `--receipts <dir>` and `--keep-specs <dir>` are opt-in.

## What it does not do (known limits)

Measured against real specs, see the compile table in `experiments/openapi-sdk`. In short:

- Only `application/json` bodies and responses; only `$ref` bodies get a typed `body` parameter; inline response schemas return untyped.
- Query parameters are all `string`; header and cookie parameters are ignored.
- `oneOf`/`anyOf`/`allOf` become `unknown` / `interface{}`; `enum` values are not kept (the alias is the base type).
- Only GET, POST, PUT, PATCH, DELETE.
- Schema `$ref`s must be `#/components/schemas/<Name>`. Any other `$ref` (e.g. a JSON pointer into `#/paths/...`, as in bundled specs) stops the run with exit 1 and one line naming the first one; nothing is written.
- No `go.mod` is written: `./sdk/go` is meant to be a package inside your own module. To check it standalone: `cd sdk/go && go mod init example.com/sdk && go build ./...`.
- Very large specs: the host (not the plugins) can run out of memory, see the table.

## Tests

`scripts/ci.sh` runs both plugins' golden tests, including `testdata/edge-input` (array/enum schemas, names with punctuation, reserved-word parameters), which compiles the output with `tsc --strict` (if `tsc` is installed) and `go vet`.

## Results on real public files

Of five public OpenAPI files we tried, one compiles as published. Measured on Linux x86_64, 1 vCPU, 2 GB RAM, about 1.2 GB free, 14 Nov 2026, on kalo-cli v0.2.0-agent.2 with recipe 0.1.1. Recipe 0.1.2 (current) changes two things: a schema named `Client` becomes `ClientModel`, and a schema that would become `ClientModel` is renamed `ClientModelModel` (#9). On 0.1.2 the Petstore output is byte-identical to `testdata/golden` (paste test, 16 Nov 2026). On 17 Nov 2026 all six rows were re-run on kalo-cli v0.2.0-agent.4 with recipe 0.1.2: same results. The Petstore output is byte-identical to `testdata/golden`, the Plaid output is byte-identical to what v0.2.0-agent.2 writes, and the other four files stop with the same message (the Stripe run with `--plugin-memory-mib 800` was not repeated). The 17 Nov re-run compared output and messages; the compile checks are from 14 Nov. Compile checks: TypeScript 7.0.2 and 5.9.3 (`tsc --strict --noEmit`) and `go vet`. Hashes are of the files as fetched on 14 Nov 2026. Linux only: no macOS or Windows run.

| File | Fetched from | sha256 | Result |
|---|---|---|---|
| Petstore (this recipe's fixture, not one of the five) | `testdata/openapi.yaml` | `cefaafa05eee75b1888762dea08a295f245e9f3e4bdb8777f4bdfed09e335339` | Compiles: `tsc --strict` and `go vet` clean. |
| Plaid 2020-09-14.yml (351 operations, 3.3 MB) | raw.githubusercontent.com/plaid/plaid-openapi/master/2020-09-14.yml | `ef42469087e2c93242baaa8b83c90fb949f44d3dbae8ac0354d4bfd4cad2947d` | Compiles: `tsc --strict` and `go vet` clean. Second run byte-identical. |
| apis.guru copy of the DigitalOcean 2.0 file | api.apis.guru/v2/specs/digitalocean.com/2.0/openapi.yaml | `2b7b040e296e92e40b5163e0bf90c3be95043ec4d27e193c491017614f69a54e` | Refused, exit 1, nothing written: its references point into `#/paths`. |
| apis.guru copy of the SendGrid file | api.apis.guru/v2/specs/sendgrid.com/1.0.0/openapi.yaml | `4c655960ff9eaf02015e1d89bbf025b08cac0b44ecbb587094e459dde5369119` | Rejected before generation, exit 1, nothing written: "decode YAML: line 13002 has control character U+009F, which YAML does not allow (3 more after it)". |
| GitHub REST api.github.com.json (14.3 MB) | raw.githubusercontent.com/github/rest-api-description/main/descriptions/api.github.com/api.github.com.json as of commit `4754642` (13 Nov 2026); `main` has changed since | `f7a3de68869181d34b232d88840d31f8157ade45cad1746eaab90cd3c801b31c` | Stopped, exit 1, nothing written: "ran out of memory at the 629 MiB plugin limit (set from free memory)"; 552 MiB on the 17 Nov run. |
| Stripe spec3.json (8.3 MB) | raw.githubusercontent.com/stripe/openapi/master/openapi/spec3.json | `7cff4cc46d0654101a36a3302f7544640f5f73e4f3cad6f264f0e23c19fa1776` | Stopped, exit 1, nothing written: "ran out of memory at the 592 MiB plugin limit (set from free memory)"; 549 MiB on the 17 Nov run. With `--plugin-memory-mib 800` it stops on "infinite circular reference detected" instead. |

The DigitalOcean and SendGrid files are apis.guru copies; the results describe the files as fetched, not the vendors' own versions. The DigitalOcean copy was the last file added, after we scanned candidates for control characters.

What the refused DigitalOcean file prints (last line):

```
SDK generation failed: unsupported $ref "#/paths/~1v2~11-clicks/get/responses/401/content/application~1json/schema" (997 such $refs in the schemas read): only $refs to #/components/schemas/<Name> are supported
```

### Memory

Plaid (3.3 MB): the whole `kalo` process peaked at 641 MiB (about 670 MB); the plugin's own memory, which the limit applies to, is part of that. The plugin memory limit is set from the memory that is free when the run starts (here about 600 MiB with about 1.2 GB free); `--plugin-memory-mib` raises it. The limit is not a fixed number: for the GitHub and Stripe files it was 629 and 592 MiB on 14 Nov (v0.2.0-agent.2) and 552 and 549 MiB on 17 Nov (v0.2.0-agent.4, about 1,050 MiB free). We have not run the GitHub or Stripe file on a larger machine. On 14 Nov: one run per file on one machine (Plaid: two; Stripe: a first run with less memory free stopped the same way at a lower limit).

### Output notes

- The TypeScript files use relative imports without a `.js` extension. `bundler` resolution is tested. An ESM package with `moduleResolution: nodenext` is not tested, and we expect it to fail.
- The output type-checks with TypeScript 7.0.2 and 5.9.3, `--strict`, on Petstore and Plaid.
- "No runtime dependency": in the Petstore golden files the TypeScript client imports only `./types`, and the Go client imports only the standard library (`bytes`, `context`, `encoding/json`, `fmt`, `io`, `net/http`, `net/url`).

### If you use openapi-typescript

openapi-typescript has been in maintenance mode since 19 Sep 2026 ([its announcement](https://github.com/openapi-ts/openapi-typescript/issues/2874)). What this recipe writes is different: interfaces and a `fetch` client class, not the `paths` type that `openapi-fetch` reads, so it does not work with `openapi-fetch`; call sites change. Unions and enum values aren't typed.
