# Contributing

Built and maintained by AI agents acting for Kalo Software GmbH. One rule: **a plugin or a recipe is a folder.**

## Add a plugin

1. `plugins/<name>/` with its own `go.mod` (module `github.com/kalo-build/kalo-plugins/plugins/<name>`), `plugin.yaml`, `cmd/plugin/main.go` (WASI entry point), tests, `testdata/` and a `README.md` with a command that works when pasted.
2. Add `use ./plugins/<name>` to `go.work`.
3. `scripts/ci.sh <name>` must pass: `go vet`, `go test`, and a reproducible `wasip1/wasm` build (`-trimpath -ldflags='-s -w -buildid='`).
4. Dependencies: MIT, Apache-2.0 or BSD only. Anything else: ask the Head of Engineering first.

Moving an existing `plugin-*` repo in: copy its tree at a tagged commit, rewrite the module path, note the source commit in the plugin README. Don't change behaviour in the same pull request.

## Add a recipe

`recipes/<name>/` holds `recipe.yaml`, `descriptors/` (SPEC v0.1-alpha descriptors), `fixtures/` (an input and the expected output) and a `README.md`. Descriptors point at wasm as `artifacts/wasm/<plugin>.wasm` with its sha256 digest and at schemas as `../../protocol/v0.1-alpha/schemas/…`. The pack step (coming with the first recipe) assembles `descriptors/ + protocol/ + artifacts/wasm/ + recipe.yaml` into `recipe-<name>-v<version>.tar.gz`; CI runs `kalo spec validate` on that assembled dir and fails if a descriptor digest doesn't match the freshly built wasm.

## Formats

`protocol/v0.1-alpha/` is a vendored, frozen copy of the SPEC schemas. Never edit it here. A change to a format id or a spec needs a conformance test and a note in the pull request.

## Review

Pull requests into `main` merge after green CI and a `LGTM` from an agent other than the author.
