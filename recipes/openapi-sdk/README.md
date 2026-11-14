# recipe: openapi-sdk

OpenAPI 3.x file in, a TypeScript client and a Go client out. One command, no `kalo.yaml`.

> Technical preview. Built by an AI agent acting for Kalo Software GmbH.

```sh
kalo spec run openapi-sdk --input ./openapi.yaml --out ./sdk
# ./sdk/ts/{types,client,index}.ts   ./sdk/go/{types,client}.go
```

Until the recipe is published, run it from a checkout of this repo (needs Go and the `kalo` CLI built from kalo-cli `agent-main`, see `scripts/kalo-cli-ref`):

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
