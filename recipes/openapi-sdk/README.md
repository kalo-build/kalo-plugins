# recipe: openapi-sdk

OpenAPI 3.x file in, a TypeScript client and a Go client out. One command, no `kalo.yaml`.

> Technical preview. Built by an AI agent acting for Kalo Software GmbH.

```sh
kalo spec run openapi-sdk --input ./openapi.yaml --out ./sdk
# ./sdk/ts/{types,client,index}.ts   ./sdk/go/{types,client}.go
```

## Install (pre-release CLI, tested on Linux)

```sh
curl -fsSL https://raw.githubusercontent.com/kalo-build/kalo-cli/agent-main/install.sh | KALO_VERSION=v0.2.0-agent.5 KALO_INSTALL_DIR="$HOME/.local/bin" sh
export PATH="$HOME/.local/bin:$PATH"
kalo spec run openapi-sdk --input ./openapi.yaml --out ./sdk
```

No root needed. The first run fetches the recipe, compiles its plugin and says so: about 15 seconds for Petstore on the machine below. After that a run takes about 2 seconds.

What it leaves on disk (one cold run, measured 25 Nov 2026, Linux x86_64, 1 vCPU, 2 GB RAM, kalo-cli v0.2.0-agent.5, recipe 0.2.0, Petstore, clean `$HOME`): the CLI prints `cached 36 MB in ~/.cache/kalo (plugins 22 MB, compiled code 14 MB)`. That is 22 MB for the recipe cache (`~/.cache/kalo/recipes`: the plugin WASM file and the descriptors) and 14 MB for the compiled plugin (`~/.cache/kalo/compiled`, which is what makes later runs take seconds), plus 15.5 MB for the binary (`~/.local/bin/kalo`). `du -sh ~/.cache/kalo` rounds to 35M, one MB below the CLI's line; both are the same 35.6 MB of files. The download is 5,590,373 bytes for the CLI tarball (`linux_x86_64`) and 4,498,169 bytes for the recipe tarball, 10.1 MB together (sizes of the release assets, not measured on the wire). Install took 0.82 s, the first run 14.18 s, the second 1.75 s. One run each, one machine. To remove all of it:

```sh
rm -rf ~/.local/bin/kalo ~/.cache/kalo
```

If `XDG_CACHE_HOME` is set, the caches are in `$XDG_CACHE_HOME/kalo` instead. `KALO_CACHE_DIR` moves them. `KALO_NO_COMPILE_CACHE=1` keeps the compiled code off disk, but then every run compiles the plugin again.

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
- Swagger 2.0 files are refused: exit 1, nothing written, and the last line says `supplied spec is a different version (oas2)`. Convert the file to OpenAPI 3.0 first.
- Schema `$ref`s must be `#/components/schemas/<Name>`. Any other `$ref` (e.g. a JSON pointer into `#/paths/...`, as in bundled specs) stops the run with exit 1 and one line naming the first one; nothing is written.
- No `go.mod` is written: `./sdk/go` is meant to be a package inside your own module. To check it standalone: `cd sdk/go && go mod init example.com/sdk && go build ./...`.
- Very large specs: the host (not the plugins) can run out of memory, see the table.

## Tests

`scripts/ci.sh` runs both plugins' golden tests, including `testdata/edge-input` (array/enum schemas, names with punctuation, reserved-word parameters), which compiles the output with `tsc --strict` (if `tsc` is installed) and `go vet`.

## Results on real public files

Of five public OpenAPI files we tried, one compiles as published. Every row below was re-run on 18 Nov 2026 on kalo-cli v0.2.0-agent.4 with recipe 0.1.3. Recipe 0.2.0 (25 Nov 2026) has not re-run these rows; its Petstore output is byte-identical to `testdata/golden`, and the 0.2.0 numbers above are from that one Petstore run. Machine: Linux x86_64, 1 vCPU, 2 GB RAM, fresh cache, default flags unless a row says otherwise. The last column gives each run's free memory at the start and, where the run printed it, the plugin memory limit. The limit is set from free memory, and a run that doesn't hit it doesn't print it. Compile checks: `tsc --strict --noEmit` with TypeScript 5.9.3 and `--moduleResolution node` (Petstore, Plaid), with TypeScript 7.0.2 and `--moduleResolution bundler` (Petstore, Plaid; TypeScript 7 removed `node`), and `go vet`. Hashes are of the files as fetched on 14 Nov 2026. Linux only: no macOS or Windows run.

History: first measured 14 Nov 2026 (v0.2.0-agent.2, recipe 0.1.1), then re-run 17 Nov 2026 (v0.2.0-agent.4, recipe 0.1.2). Each re-run gave the same results, except the Stripe row (see that row). Recipe 0.1.2 changed two things: a schema named `Client` becomes `ClientModel`, and a schema that would become `ClientModel` is renamed `ClientModelModel` (#9). Recipe 0.1.3 changed two more, in both languages. A 200/201 response declared as an inline array of a `$ref` returns the typed list (`Promise<Note[]>`, `([]Note, error)`). A response that wraps a `$ref`, or an array of one, in a `data` field returns that field instead of the whole body. Neither touches the Petstore output, which stays byte-identical to `testdata/golden` (CI, 18 Nov 2026).

| File | Fetched from | sha256 | Result | Run (18 Nov 2026, v0.2.0-agent.4, recipe 0.1.3) |
|---|---|---|---|---|
| Petstore (this recipe's fixture, not one of the five) | `testdata/openapi.yaml` | `cefaafa05eee75b1888762dea08a295f245e9f3e4bdb8777f4bdfed09e335339` | Compiles: `tsc --strict` and `go vet` clean. | Free memory not logged, peak 343 MiB, limit not reached. |
| Plaid 2020-09-14.yml (351 operations, 3.3 MB) | raw.githubusercontent.com/plaid/plaid-openapi/master/2020-09-14.yml | `ef42469087e2c93242baaa8b83c90fb949f44d3dbae8ac0354d4bfd4cad2947d` | Compiles: `tsc --strict` and `go vet` clean. Second run byte-identical. | Three runs: 1,281 MiB free on the last, peak 563 MiB (533 and 567 MiB on the other two), limit not reached. |
| apis.guru copy of the DigitalOcean 2.0 file | api.apis.guru/v2/specs/digitalocean.com/2.0/openapi.yaml | `2b7b040e296e92e40b5163e0bf90c3be95043ec4d27e193c491017614f69a54e` | Refused, exit 1, nothing written: its references point into `#/paths`. | 1,147 MiB free, peak 272 MiB, stops before the limit matters. |
| apis.guru copy of the SendGrid file | api.apis.guru/v2/specs/sendgrid.com/1.0.0/openapi.yaml | `4c655960ff9eaf02015e1d89bbf025b08cac0b44ecbb587094e459dde5369119` | Rejected before generation, exit 1, nothing written: "decode YAML: line 13002 has control character U+009F, which YAML does not allow (3 more after it)". | 1,208 MiB free, peak 59 MiB, stops before the limit matters. |
| GitHub REST api.github.com.json (14.3 MB) | raw.githubusercontent.com/github/rest-api-description/main/descriptions/api.github.com/api.github.com.json as of commit `4754642` (13 Nov 2026); `main` has changed since | `f7a3de68869181d34b232d88840d31f8157ade45cad1746eaab90cd3c801b31c` | Stopped, exit 1, nothing written: "ran out of memory at the 746 MiB plugin limit (set from free memory)" on 18 Nov. Same stop at 629 MiB (14 Nov) and 552 MiB (17 Nov). | 1,231 MiB free, **746 MiB limit** (out of memory), peak 846 MiB. |
| Stripe spec3.json (8.3 MB) | raw.githubusercontent.com/stripe/openapi/master/openapi/spec3.json | `7cff4cc46d0654101a36a3302f7544640f5f73e4f3cad6f264f0e23c19fa1776` | Stopped, exit 1, nothing written. With about 1,270 MiB free (18 Nov) it stops on "infinite circular reference detected: payment_intent: payment_intent -> api_errors -> payment_intent", with default flags and again with `--plugin-memory-mib 800`. With less free memory it runs out of memory first: at a 592 MiB limit on 14 Nov and 549 MiB on 17 Nov (about 1,050 MiB free). | 1,267 MiB free, default flags, peak 883 MiB: circular reference, not out of memory. |

The DigitalOcean and SendGrid files are apis.guru copies; the results describe the files as fetched, not the vendors' own versions. The DigitalOcean copy was the last file added, after we scanned candidates for control characters.

What the refused DigitalOcean file prints (last line):

```
SDK generation failed: unsupported $ref "#/paths/~1v2~11-clicks/get/responses/401/content/application~1json/schema" (997 such $refs in the schemas read): only $refs to #/components/schemas/<Name> are supported
```

### Memory

Plaid (3.3 MB): the whole `kalo` process peaked at 641 MiB (about 670 MB); the plugin's own memory, which the limit applies to, is part of that. The plugin memory limit is set from the memory that is free when the run starts (here about 600 MiB with about 1.2 GB free); `--plugin-memory-mib` raises it. The limit is not a fixed number: for the GitHub and Stripe files it was 629 and 592 MiB on 14 Nov (v0.2.0-agent.2) and 552 and 549 MiB on 17 Nov (v0.2.0-agent.4, about 1,050 MiB free). On 18 Nov (v0.2.0-agent.4, recipe 0.1.3, about 1,230 to 1,280 MiB free) GitHub stopped at a 746 MiB limit. Stripe did not reach its limit and stopped on the circular reference. We have not run the GitHub or Stripe file on a larger machine. On 14 Nov: one run per file on one machine (Plaid: two; Stripe: a first run with less memory free stopped the same way at a lower limit).

### Output notes

- The TypeScript files use relative imports without a `.js` extension. `bundler` resolution is tested. An ESM package with `moduleResolution: nodenext` is not tested, and we expect it to fail.
- The output type-checks with TypeScript 7.0.2 and 5.9.3, `--strict`, on Petstore and Plaid.
- "No runtime dependency": in the Petstore golden files the TypeScript client imports only `./types`, and the Go client imports only the standard library (`bytes`, `context`, `encoding/json`, `fmt`, `io`, `net/http`, `net/url`).

### If you use openapi-typescript

openapi-typescript has been in maintenance mode since 19 Sep 2026 ([its announcement](https://github.com/openapi-ts/openapi-typescript/issues/2874)). What this recipe writes is different: interfaces and a `fetch` client class, not the `paths` type that `openapi-fetch` reads, so it does not work with `openapi-fetch`; call sites change. Unions and enum values aren't typed.
