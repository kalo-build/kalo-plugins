# Task hq-78

## Task
One-wasm openapi-sdk recipe (0.2.0): build to LGTM against bars; no merge, no tag

## Status
packed

## Series
D

## Profile
code-change

## Created At
2026-11-19T09:05:03Z

## Original Query
One-wasm openapi-sdk recipe (0.2.0): build to LGTM against bars; no merge, no tag

## Repo / Workspace
- Repo: `/data/workspaces/eng_head/kalo-plugins`
- Workspace: `/data/workspaces/eng_head/kalo-plugins/devspecs/tasks/hq-78`

## Resources
- `task.json`
- `D01-one-wasm-openapi-sdk-recipe-0-2-0-build-to-lgtm-against-bars-no-plan.md`
- `D01-one-wasm-openapi-sdk-recipe-0-2-0-build-to-lgtm-against-bars-no-result.md`

## Task Slices
- D01: One-wasm openapi-sdk recipe (0.2.0): build to LGTM against bars; no merge, no tag. Plan: `D01-one-wasm-openapi-sdk-recipe-0-2-0-build-to-lgtm-against-bars-no-plan.md`. Result: `D01-one-wasm-openapi-sdk-recipe-0-2-0-build-to-lgtm-against-bars-no-result.md`.

## Relevant Map Areas
- `plugins`
- `scripts`

## Likely Primary Files
- `scripts/stage-recipe.sh`
  Evidence: anchor-first ranking: score 24.000; matches openapi-sdk, openapisdk, recipe; fields body, path, title; query term match in path: recipe; query term match in path: tag
- `plugins/openapi-sdk-ts/pkg/generate/client.go` - plugins/openapi-sdk-ts/pkg/generate/client.go (go)
  Evidence: query term match in path: openapi; query term match in path: openapi-sdk; query term match in path: sdk
- `plugins/openapi-sdk-ts/pkg/generate/generate.go` - plugins/openapi-sdk-ts/pkg/generate/generate.go (go)
  Evidence: query term match in path: openapi; query term match in path: openapi-sdk; query term match in path: sdk
- `plugins/openapi-sdk-go/pkg/generate/client.go`
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-go/pkg/generate/types.go`
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-go/pkg/generate/generate.go` - plugins/openapi-sdk-go/pkg/generate/generate.go (go)
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk

## Likely Tests
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: openapi
- `plugins/openapi-sdk-go/pkg/generate/linear_test.go`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: openapi
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go#L14` - TestUnsupportedRef_FailsWithOneLine
  Evidence: relationship expansion: source_manifest_relationship_recovery; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go#L14` - TestUnsupportedRef_FailsWithOneLine
  Evidence: relationship expansion: source_manifest_relationship_recovery; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk
- `scripts/test-recipe.sh`
  Evidence: relationship expansion: source_manifest_relationship_recovery; query term match in path: recipe; query term match in body: against
- `plugins/openapi-sdk-go/pkg/generate/selfref_test.go#L32` - TestSelfRef_OptionalDirect
  Evidence: relationship expansion: source_manifest_relationship_recovery; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk

## Likely Docs / Plans / Config
None found in the initial preflight.

## Supporting Context
None found in the initial preflight.

## Related Git Receipts
- `4573f58` 2026-11-14 - openapi-sdk-ts/go: type aliases for non-object schemas, identifier sanitising, edge fixture + goldens
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/client.go`, `plugins/openapi-sdk-go/pkg/generate/edge_test.go`, `plugins/openapi-sdk-go/pkg/generate/types.go`
- `d99d1b4` 2026-11-14 - chore: monorepo layout; move plugin-openapi-sdk-ts in as plugins/openapi-sdk-ts (#1)
  Matched paths: `plugins/openapi-sdk-ts/pkg/generate/client.go`, `plugins/openapi-sdk-ts/pkg/generate/generate.go`, `plugins/openapi-sdk-ts/pkg/generate/types.go`
- `c5e1c0f` 2026-11-18 - openapi-sdk-ts/go: optional self-reference generates instead of recursing (hq-74)
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/generate.go`, `plugins/openapi-sdk-ts/pkg/generate/generate.go`

## Noise Risks
- `recipes/openapi-sdk/testdata/golden/ts/types.ts` - recipes/openapi-sdk/testdata/golden/ts/types.ts (typescript)
  Evidence: anchor-first ranking: score 24.000; matches openapi-sdk, openapisdk; fields path; query term match in path: openapi; query term match in path: openapi-sdk
- `recipes/openapi-sdk/testdata/golden/ts/client.ts` - recipes/openapi-sdk/testdata/golden/ts/client.ts (typescript)
  Evidence: anchor-first ranking: score 24.000; matches openapi-sdk, openapisdk; fields path; query term match in path: openapi; query term match in path: openapi-sdk
- `recipes/openapi-sdk/testdata/golden/ts/index.ts` - recipes/openapi-sdk/testdata/golden/ts/index.ts (typescript)
  Evidence: anchor-first ranking: score 24.000; matches openapi-sdk, openapisdk; fields path; query term match in path: openapi; query term match in path: openapi-sdk

## Freshness Warnings
These on-disk paths match the task wording but were not present in the indexed candidate set. Treat them as stale-index risk, not proof that the initial pack is wrong.

- `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe
- `recipes/openapi-sdk/descriptors/bindings/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe
- `recipes/openapi-sdk/descriptors/bindings/openapi.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe
- `recipes/openapi-sdk/descriptors/bindings/ts-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe
- `recipes/openapi-sdk/descriptors/contracts/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe
- `recipes/openapi-sdk/descriptors/contracts/openapi.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe
- `recipes/openapi-sdk/descriptors/contracts/ts-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe
- `recipes/openapi-sdk/descriptors/implementations/go.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe

## Risk Cards
Evidence-backed checks to run before trusting the initial task context. These are not required edit targets.

- On-disk paths matched the task but were not indexed [medium, freshness]
  Agent check: Inspect the warned files or refresh the index before trusting missing context.
  Evidence: `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe; `recipes/openapi-sdk/descriptors/bindings/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe; `recipes/openapi-sdk/descriptors/bindings/openapi.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe; `recipes/openapi-sdk/descriptors/bindings/ts-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe

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
- found 7 likely primary file(s)
- found 8 likely test file(s)
- found 3 related Git receipt(s)
- 3 file(s) were downgraded as likely noise

Agent instruction:
Validate the test and integration surface before editing. Record critical misses and distracting inclusions in the slice result or a task checkpoint.

## Suggested Starting Slice
Use `D01-one-wasm-openapi-sdk-recipe-0-2-0-build-to-lgtm-against-bars-no-plan.md` as the first bounded plan in this task thread. Refine it before editing if primary files, tests, or integration points look incomplete.

## Agent Preflight Checklist
- [ ] Verify the likely primary files against the repo before editing.
- [ ] Search for same-package or same-command tests if test confidence is not high.
- [ ] Check receipt-touched related files before assuming the pack is complete.
- [ ] Record files actually read, edited, tests run, misses, and noise in `D01-one-wasm-openapi-sdk-recipe-0-2-0-build-to-lgtm-against-bars-no-result.md` or `ds task checkpoint`.
