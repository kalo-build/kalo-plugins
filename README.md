# kalo-plugins

Plugins and one-shot recipes for [Kalo](https://kalo.build), the open-source CLI that runs pipelines of WASM plugins. **Technical preview.**

This repository is built and maintained by AI agents acting for Kalo Software GmbH.

## Quick start: OpenAPI file in, TypeScript and Go SDKs out

```sh
curl -fsSL https://raw.githubusercontent.com/kalo-build/kalo-cli/agent-main/install.sh | KALO_VERSION=v0.2.0-agent.1 sh
kalo spec run openapi-sdk --input ./openapi.yaml --out ./sdk
```

Writes `./sdk/ts/` and `./sdk/go/` and nothing else. Runs locally, no account. Pre-release CLI build, tested on Linux. Limits and the specs it was checked against: [`recipes/openapi-sdk`](recipes/openapi-sdk/README.md).

## Layout

```
plugins/<name>/          one plugin = one folder (own go.mod, plugin.yaml, tests)
recipes/<name>/          one recipe = one folder (SPEC descriptors + recipe.yaml + fixtures)
protocol/v0.1-alpha/     vendored SPEC v0.1-alpha schemas (frozen, do not edit)
scripts/ci.sh            builds every plugin to WASM and runs every test
```

## Build and test everything

Needs Go 1.24+.

```sh
git clone https://github.com/kalo-build/kalo-plugins && cd kalo-plugins
scripts/ci.sh
```

WASM files land in `dist/wasm/`.

## Plugins

| Plugin | What it does | From |
|---|---|---|
| [`openapi-sdk-ts`](plugins/openapi-sdk-ts) | OpenAPI 3.1 YAML → typed TypeScript HTTP client | `kalo-build/plugin-openapi-sdk-ts@7d36150` |
| [`openapi-sdk-go`](plugins/openapi-sdk-go) | OpenAPI 3.1 YAML → typed Go HTTP client | `kalo-build/plugin-openapi-sdk-go@d177032` |

Adding a plugin or a recipe: [CONTRIBUTING.md](CONTRIBUTING.md). Licence: MIT.
