# Task hq-74

## Task
Recipe 0.1.4: optional self-reference generates (openapi-sdk-ts/go), PR only, no tag before the website merge

## Status
packed

## Series
C

## Profile
code-change

## Created At
2026-11-18T18:35:16Z

## Original Query
Recipe 0.1.4: optional self-reference generates (openapi-sdk-ts/go), PR only, no tag before the website merge

## Repo / Workspace
- Repo: `/data/workspaces/eng_head/kalo-plugins`
- Workspace: `/data/workspaces/eng_head/kalo-plugins/devspecs/tasks/hq-74`

## Resources
- `task.json`
- `C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-plan.md`
- `C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-result.md`

## Task Slices
- C01: Recipe 0.1.4: optional self-reference generates (openapi-sdk-ts/go), PR only, no tag before the website merge. Plan: `C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-plan.md`. Result: `C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-result.md`.

## Relevant Map Areas
- `plugins`
- `scripts`

## Likely Primary Files
- `plugins/openapi-sdk-ts/pkg/generate/generate.go` - plugins/openapi-sdk-ts/pkg/generate/generate.go (go)
  Evidence: query term match in path: openapi; query term match in path: openapi-sdk-ts; query term match in path: sdk
- `plugins/openapi-sdk-go/pkg/generate/generate.go` - plugins/openapi-sdk-go/pkg/generate/generate.go (go)
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: sdk
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: openapi-sdk-ts
- `scripts/stage-recipe.sh`
  Evidence: anchor-first ranking: score 24.000; matches recipe; fields path, title, body; query term match in path: recipe; query term match in path: tag
- `plugins/openapi-sdk-ts/pkg/generate/warn.go` - plugins/openapi-sdk-ts/pkg/generate/warn.go (go)
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk-ts

## Likely Tests
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L57` - TestGenerateSDK_TypesOnly
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; anchor-first ranking: score 24.000; matches only; fields title, test_name, symbol, body; query term match in path: openapi
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L80` - TestGenerateSDK_ClientOnly
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; anchor-first ranking: score 24.000; matches only; fields title, test_name, symbol, body; query term match in path: openapi
- `plugins/openapi-sdk-ts/pkg/generate/notes_test.go`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: openapi
- `plugins/openapi-sdk-ts/pkg/generate/shapes_test.go`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: openapi
- `plugins/openapi-sdk-ts/pkg/generate/compile_test.go#L21` - TestCompileTestSuite
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: openapi-sdk-ts
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: sdk
- `scripts/test-recipe.sh`
  Evidence: relationship expansion: source_manifest_relationship_recovery; query term match in path: recipe; query term match in body: openapi
- `plugins/openapi-sdk-ts/pkg/generate/warn_test.go#L8` - TestWarnIfOpenAPI31
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk-ts

## Likely Docs / Plans / Config
None found in the initial preflight.

## Supporting Context
None found in the initial preflight.

## Related Git Receipts
- `b6ea57a` 2026-11-18 - openapi-sdk-ts/go: inline array of $ref typed, data field unwrapped; recipe 0.1.3 (hq-68) (#17)
  Matched paths: `plugins/openapi-sdk-ts/pkg/generate/notes_test.go`, `plugins/openapi-sdk-ts/pkg/generate/shapes_test.go`
- `5d39ace` 2026-11-14 - openapi-sdk-ts/go: ApiError clash, required-before-optional params, only used Go imports, OpenAPI 3.1 warning (experi...
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/generate.go`, `plugins/openapi-sdk-ts/pkg/generate/generate.go`, `plugins/openapi-sdk-ts/pkg/generate/shapes_test.go`
- `d99d1b4` 2026-11-14 - chore: monorepo layout; move plugin-openapi-sdk-ts in as plugins/openapi-sdk-ts (#1)
  Matched paths: `plugins/openapi-sdk-ts/pkg/generate/generate.go`, `plugins/openapi-sdk-ts/pkg/generate/generate_test.go`, `plugins/openapi-sdk-ts/pkg/generate/types.go`

## Noise Risks
None found in the initial preflight.

## Freshness Warnings
These on-disk paths match the task wording but were not present in the indexed candidate set. Treat them as stale-index risk, not proof that the initial pack is wrong.

- `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk
- `recipes/openapi-sdk/descriptors/bindings/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk
- `recipes/openapi-sdk/descriptors/bindings/openapi.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk
- `recipes/openapi-sdk/descriptors/bindings/ts-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk
- `recipes/openapi-sdk/descriptors/contracts/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk
- `recipes/openapi-sdk/descriptors/contracts/openapi.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk
- `recipes/openapi-sdk/descriptors/contracts/ts-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk
- `recipes/openapi-sdk/descriptors/implementations/go.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk

## Risk Cards
Evidence-backed checks to run before trusting the initial task context. These are not required edit targets.

- On-disk paths matched the task but were not indexed [medium, freshness]
  Agent check: Inspect the warned files or refresh the index before trusting missing context.
  Evidence: `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk; `recipes/openapi-sdk/descriptors/bindings/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk; `recipes/openapi-sdk/descriptors/bindings/openapi.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk; `recipes/openapi-sdk/descriptors/bindings/ts-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk

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
- Noise risk: low
- Pack completeness: medium

Why:
- found 5 likely primary file(s)
- found 8 likely test file(s)
- found 3 related Git receipt(s)

Agent instruction:
Validate the test and integration surface before editing. Record critical misses and distracting inclusions in the slice result or a task checkpoint.

## Suggested Starting Slice
Use `C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-plan.md` as the first bounded plan in this task thread. Refine it before editing if primary files, tests, or integration points look incomplete.

## Agent Preflight Checklist
- [ ] Verify the likely primary files against the repo before editing.
- [ ] Search for same-package or same-command tests if test confidence is not high.
- [ ] Check receipt-touched related files before assuming the pack is complete.
- [ ] Record files actually read, edited, tests run, misses, and noise in `C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-result.md` or `ds task checkpoint`.
