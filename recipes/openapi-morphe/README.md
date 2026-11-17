# openapi-morphe recipe

**Technical preview.** OpenAPI file in, Morphe structures (`.str`) and enums (`.enum`) out, with one command and no `kalo.yaml`.

## Run it

Install the pinned preview build of `kalo` (the same install line as the openapi-sdk recipe):

```sh
curl -fsSL https://raw.githubusercontent.com/kalo-build/kalo-cli/agent-main/install.sh | KALO_VERSION=v0.2.0-agent.4 KALO_INSTALL_DIR="$HOME/.local/bin" sh
```

Then, next to your OpenAPI file:

```sh
kalo spec run openapi-morphe --input ./openapi.yaml --out ./specs
```

`kalo` downloads the release asset `recipe-openapi-morphe-v<version>.tar.gz` from this repository, checks its sha256 and the plugin's digest, and runs the plugin locally. Your file is not uploaded anywhere.

## What comes out

Files under `./specs/morphe/`:

- `structures/<name>.str` for each object schema in `components.schemas`: fields with Morphe types.
- A schema that is neither an object nor a string enum (for example an array) produces no file and no message.
- `enums/<name>.enum` for string enums, including string enums inlined in a property (named after schema and property, e.g. `payment_create_method.enum`).

For the fixture in `testdata/openapi.yaml` the output is exactly `testdata/golden/morphe/` (9 files); CI checks that on every pull request.

## What it does not do

- It reads `components.schemas` only. Paths, operations, parameters and security schemes are not turned into Morphe.
- It writes structures and enums, not Morphe models or entities.
- It does not generate code from the Morphe files. The Morphe plugins that take this folder (for example to TypeScript types or Go structs) run through a `kalo.yaml` today; see the [kalo-cli README](https://github.com/kalo-build/kalo-cli#readme). There is no recipe for that step yet.

## Results on public files

Released `kalo` v0.2.0-agent.4, default flags, 2 GB machine (about 830 MiB free at the start), one run per file, recipe staged from this folder with `--recipe-dir` on 17 Nov 2026. Same six files as the openapi-sdk recipe. "Peak" is the whole `kalo` process.

| File | Result |
|---|---|
| Swagger Petstore (2.7 KB) | Exit 0, 2 files (`pet.str`, `error.str`), 11 s with the first-run plugin compile, peak 321 MiB. The `Pets` schema (an array of `Pet`) produces no file and no message. |
| Plaid 2020-09-14.yml (3.3 MB) | Stopped, exit 1, nothing written: "ran out of memory at the 288 MiB plugin limit (set from free memory)". |
| apis.guru copy of the DigitalOcean file (1.6 MB) | Exit 1, nothing written: "plugin wrote no files". The copy has no `components.schemas` block, so there is nothing to convert; the message does not say that. |
| apis.guru copy of the SendGrid file | Rejected before the plugin runs, exit 1, nothing written: "line 13002 has control character U+009F, which YAML does not allow". |
| GitHub REST api.github.com.json (14.3 MB) | Stopped, exit 1, nothing written: "ran out of memory at the 419 MiB plugin limit (set from free memory)". |
| Stripe spec3.json (8.3 MB) | Stopped, exit 1, nothing written: "ran out of memory at the 358 MiB plugin limit (set from free memory)". |

The plugin memory limit is set from the memory that is free when the run starts; `--plugin-memory-mib` raises it. We have not run the large files with more memory.

## The plugin

The recipe runs `@kalo-build/plugin-openapi-morphe` v1.0.0 exactly as published on registry.kalo.build (sha256 `35379df2f6f09895a91f5be92bb29d63ac7b14852af95d80a2285e612d0329c1`, pinned in `artifacts.lock`). Source: [kalo-build/plugin-openapi-morphe](https://github.com/kalo-build/plugin-openapi-morphe) at `fcde75c`, MIT licence; its notice ships in the tarball as `third-party/openapi-morphe-LICENSE`. This folder holds no copy of the plugin's source.

---

Built and written by an AI agent (eng_head) acting for Kalo Software GmbH.
