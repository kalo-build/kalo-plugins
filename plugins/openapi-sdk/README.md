# openapi-sdk (one wasm for the openapi-sdk recipe)

Both openapi-sdk generators, [`openapi-sdk-ts`](../openapi-sdk-ts) and [`openapi-sdk-go`](../openapi-sdk-go), linked into one WASI binary. The config field `target` picks the generator: `ts` or `go`. Everything else (config fields, exit codes, messages, output) is what the two single-target plugins do; this folder holds only the entry point, the generators stay in their own folders.

The [openapi-sdk recipe](../../recipes/openapi-sdk) (0.2.0 and later) ships this one artifact. Its two processor implementations (`KALOX:OPENAPI-TO-TS-SDK@1`, `KALOX:OPENAPI-TO-GO-SDK@1`) point at the same file and digest, each with its own fixed `target`. Why: libopenapi, the Go runtime and the YAML code are shared, so one binary is about the size of one of the two (21.8 MB instead of 42.0 MB raw), and the CLI compiles and caches it once.

It is not a registry plugin (no `plugin.yaml`); the registry keeps the two single-target plugins.

## Build and test

```sh
scripts/ci.sh openapi-sdk          # go vet, go test, then dist/wasm/openapi-sdk.wasm
```

The test runs both targets on the recipe fixture and compares every file with `recipes/openapi-sdk/testdata/golden`.

## Config

```json
{"inputPath": "/input", "outputPath": "/output",
 "config": {"target": "ts", "specFileName": "openapi.yaml"}}
```

`target` is required: a missing or unknown value exits 4 and writes nothing. `ts` also reads `typesImportPath`; `go` reads `packageName` and `packagePath`; both read `generateTypes` and `generateClient`.

Technical preview, built by an AI agent acting for Kalo Software GmbH.
