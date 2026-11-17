# Task hq-52

## Task
kalo-plugins README + recipe README: pin agent.2, OpenAPI 3.0 wording, drop the stale "until published" block

## Status
packed

## Series
A

## Profile
code-change

## Created At
2026-11-17T15:40:14Z

## Original Query
kalo-plugins README + recipe README: pin agent.2, OpenAPI 3.0 wording, drop the stale "until published" block

## Repo / Workspace
- Repo: `/tmp/kp`
- Workspace: `/tmp/kp/devspecs/tasks/hq-52`

## Resources
- `task.json`
- `A01-kalo-plugins-readme-recipe-readme-pin-agent-2-openapi-3-0-plan.md`
- `A01-kalo-plugins-readme-recipe-readme-pin-agent-2-openapi-3-0-result.md`

## Task Slices
- A01: kalo-plugins README + recipe README: pin agent.2, OpenAPI 3.0 wording, drop the stale "until published" block. Plan: `A01-kalo-plugins-readme-recipe-readme-pin-agent-2-openapi-3-0-plan.md`. Result: `A01-kalo-plugins-readme-recipe-readme-pin-agent-2-openapi-3-0-result.md`.

## Relevant Map Areas
- `plugins`
- `scripts`

## Likely Primary Files
- `plugins/openapi-sdk-go/pkg/generate/generate.go` - plugins/openapi-sdk-go/pkg/generate/generate.go (go)
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/refs.go` - plugins/openapi-sdk-go/pkg/generate/refs.go (go)
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/spec.go` - plugins/openapi-sdk-go/pkg/generate/spec.go (go)
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/warn.go` - plugins/openapi-sdk-go/pkg/generate/warn.go (go)
  Evidence: query term match in path: openapi; query term match in path: plugins; identifier/body match: 3.0
- `plugins/openapi-sdk-ts/pkg/generate/generate.go` - plugins/openapi-sdk-ts/pkg/generate/generate.go (go)
  Evidence: query term match in path: openapi; query term match in path: plugins; scout source preservation: query roots opena,openapi,plugi,plugin; primary test roots gener,genera,generate; top source rank 5
- `plugins/openapi-sdk-ts/pkg/generate/refs.go` - plugins/openapi-sdk-ts/pkg/generate/refs.go (go)
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/spec.go` - plugins/openapi-sdk-ts/pkg/generate/spec.go (go)
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/warn.go` - plugins/openapi-sdk-ts/pkg/generate/warn.go (go)
  Evidence: query term match in path: openapi; query term match in path: plugins; identifier/body match: 3.0
- `scripts/pack-recipe.sh`
  Evidence: query term match in path: recipe; query term match in body: kalo; query term match in body: openapi
- `scripts/stage-recipe.sh`
  Evidence: query term match in path: recipe; query term match in body: openapi; query term match in body: plugins
- `plugins/openapi-sdk-go/pkg/generate/config.go`
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/config.go`
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/cmd/plugin/main.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: kalo
- `plugins/openapi-sdk-go/pkg/generate/types.go`
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/cmd/plugin/main.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: kalo
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/util.go`
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/client.go`
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/util.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: pin
- `plugins/openapi-sdk-ts/pkg/generate/client.go`
  Evidence: query term match in path: openapi; query term match in path: plugins

## Likely Tests
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L61` - TestGenerateSDK_TypesOnly
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/warn_test.go#L8` - TestWarnIfOpenAPI31
  Evidence: query term match in path: openapi; query term match in path: plugins; identifier/body match: 3.0
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: kalo
- `plugins/openapi-sdk-go/pkg/generate/compile_test.go#L21` - TestCompileTestSuite
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go#L18` - TestEdgeSpec_Golden
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L86` - TestParseSpec
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L99` - TestGenerateTypes_Schemas
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L118` - TestGenerateClient_Operations
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L137` - TestConfig_Resolve
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go#L14` - TestUnsupportedRef_FailsWithOneLine
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/shapes_test.go#L18` - TestShapesSpec_Golden
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/compile_test.go#L21` - TestCompileTestSuite
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/edge_test.go#L18` - TestEdgeSpec_Golden
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L57` - TestGenerateSDK_TypesOnly
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L80` - TestGenerateSDK_ClientOnly
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L103` - TestGenerateSDK_ExternalTypes
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L131` - TestParseSpec
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L144` - TestGenerateTypes_Schemas
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L161` - TestGenerateClient_Operations
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L178` - TestConfig_Resolve
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L185` - TestConfig_Resolve_PreservesExplicit
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go#L14` - TestUnsupportedRef_FailsWithOneLine
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/shapes_test.go#L18` - TestShapesSpec_Golden
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/warn_test.go#L8` - TestWarnIfOpenAPI31
  Evidence: query term match in path: openapi; query term match in path: plugins; identifier/body match: 3.0
- `scripts/test-recipe.sh`
  Evidence: relationship expansion: source_manifest_relationship_recovery; query term match in path: recipe; query term match in body: kalo
- `plugins/openapi-sdk-go/pkg/generate/warn_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; identifier/body match: 3.0
- `plugins/openapi-sdk-ts/pkg/generate/warn_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; identifier/body match: 3.0
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: kalo
- `plugins/openapi-sdk-go/pkg/generate/shapes_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: kalo
- `plugins/openapi-sdk-ts/pkg/generate/edge_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: kalo
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: kalo
- `plugins/openapi-sdk-ts/pkg/generate/shapes_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: kalo
- `plugins/openapi-sdk-go/pkg/generate/compile_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: kalo
- `plugins/openapi-sdk-ts/pkg/generate/compile_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: kalo
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins

## Likely Docs / Plans / Config
None found in the initial preflight.

## Supporting Context
None found in the initial preflight.

## Related Git Receipts
- `f5ffed5` 2026-11-14 - openapi-sdk-ts/go: refuse schema $refs outside #/components/schemas with one line and a non-zero exit (experiment ope...
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/generate.go`, `plugins/openapi-sdk-go/pkg/generate/refs.go`, `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `5d39ace` 2026-11-14 - openapi-sdk-ts/go: ApiError clash, required-before-optional params, only used Go imports, OpenAPI 3.1 warning (experi...
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/generate.go`, `plugins/openapi-sdk-go/pkg/generate/warn.go`, `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `d99d1b4` 2026-11-14 - chore: monorepo layout; move plugin-openapi-sdk-ts in as plugins/openapi-sdk-ts (#1)
  Matched paths: `plugins/openapi-sdk-ts/pkg/generate/generate.go`, `plugins/openapi-sdk-ts/pkg/generate/spec.go`

## Noise Risks
- `plugins/openapi-sdk-ts/testdata/edge-ground-truth/client.ts` - plugins/openapi-sdk-ts/testdata/edge-ground-truth/client.ts (typescript)
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/testdata/edge-ground-truth/types.ts` - plugins/openapi-sdk-ts/testdata/edge-ground-truth/types.ts (typescript)
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/testdata/ground-truth/client.ts` - plugins/openapi-sdk-ts/testdata/ground-truth/client.ts (typescript)
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/testdata/ground-truth/types.ts` - plugins/openapi-sdk-ts/testdata/ground-truth/types.ts (typescript)
  Evidence: query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/testdata/shapes-ground-truth/client.ts` - plugins/openapi-sdk-ts/testdata/shapes-ground-truth/client.ts (typescript)
  Evidence: query term match in path: openapi; query term match in path: plugins

## Freshness Warnings
These on-disk paths match the task wording but were not present in the indexed candidate set. Treat them as stale-index risk, not proof that the initial pack is wrong.

- `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: plugins, readme, openapi
- `plugins/openapi-sdk-ts/README.md` - on-disk path matched task terms but was not in the indexed candidate set: plugins, readme, openapi
- `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: readme, recipe, openapi
- `recipes/openapi-sdk/descriptors/namespaces/kalox.yaml` - on-disk path matched task terms but was not in the indexed candidate set: kalo, recipe, openapi
- `plugins/openapi-sdk-go/plugin.yaml` - on-disk path matched task terms but was not in the indexed candidate set: plugins, openapi
- `plugins/openapi-sdk-ts/plugin.yaml` - on-disk path matched task terms but was not in the indexed candidate set: plugins, openapi
- `recipes/openapi-sdk/descriptors/bindings/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi
- `recipes/openapi-sdk/descriptors/bindings/openapi.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi

## Risk Cards
Evidence-backed checks to run before trusting the initial task context. These are not required edit targets.

- On-disk paths matched the task but were not indexed [medium, freshness]
  Agent check: Inspect the warned files or refresh the index before trusting missing context.
  Evidence: `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: plugins, readme, openapi; `plugins/openapi-sdk-ts/README.md` - on-disk path matched task terms but was not in the indexed candidate set: plugins, readme, openapi; `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: readme, recipe, openapi; `recipes/openapi-sdk/descriptors/namespaces/kalox.yaml` - on-disk path matched task terms but was not in the indexed candidate set: kalo, recipe, openapi

## Known Knowns
- The preflight found likely primary implementation files.
- The preflight found likely behavior/test artifacts.
- Git receipts provide historical trust evidence for packed paths.

## Known Unknowns
- Task-related on-disk paths may be missing from the indexed candidate set.
- Pack completeness is not high; verify the working set before editing.

## Confidence Summary
- Primary file confidence: high
- Test coverage confidence: high
- Docs/config coverage confidence: low
- Git receipt confidence: high
- Noise risk: high
- Pack completeness: low

Why:
- found 20 likely primary file(s)
- found 38 likely test file(s)
- found 3 related Git receipt(s)
- 5 file(s) were downgraded as likely noise

Agent instruction:
Validate the test and integration surface before editing. Record critical misses and distracting inclusions in the slice result or a task checkpoint.

## Suggested Starting Slice
Use `A01-kalo-plugins-readme-recipe-readme-pin-agent-2-openapi-3-0-plan.md` as the first bounded plan in this task thread. Refine it before editing if primary files, tests, or integration points look incomplete.

## Agent Preflight Checklist
- [ ] Verify the likely primary files against the repo before editing.
- [ ] Search for same-package or same-command tests if test confidence is not high.
- [ ] Check receipt-touched related files before assuming the pack is complete.
- [ ] Record files actually read, edited, tests run, misses, and noise in `A01-kalo-plugins-readme-recipe-readme-pin-agent-2-openapi-3-0-result.md` or `ds task checkpoint`.
