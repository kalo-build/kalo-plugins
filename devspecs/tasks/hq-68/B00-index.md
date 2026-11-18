# Task hq-68

## Task
Recipe openapi-sdk 0.1.3: a 200/201 response that is an inline array of $ref is typed in TS and Go (kalo-plugins only, stop before the tag)

## Status
packed

## Series
B

## Profile
code-change

## Created At
2026-11-18T15:32:19Z

## Original Query
Recipe openapi-sdk 0.1.3: a 200/201 response that is an inline array of $ref is typed in TS and Go (kalo-plugins only, stop before the tag)

## Repo / Workspace
- Repo: `/data/workspaces/eng_head/kalo-plugins`
- Workspace: `/data/workspaces/eng_head/kalo-plugins/devspecs/tasks/hq-68`

## Resources
- `task.json`
- `B01-recipe-openapi-sdk-0-1-3-a-200-201-response-that-is-an-inline-plan.md`
- `B01-recipe-openapi-sdk-0-1-3-a-200-201-response-that-is-an-inline-result.md`

## Task Slices
- B01: Recipe openapi-sdk 0.1.3: a 200/201 response that is an inline array of $ref is typed in TS and Go (kalo-plugins only, stop before the tag). Plan: `B01-recipe-openapi-sdk-0-1-3-a-200-201-response-that-is-an-inline-plan.md`. Result: `B01-recipe-openapi-sdk-0-1-3-a-200-201-response-that-is-an-inline-result.md`.

## Relevant Map Areas
- `plugins`
- `scripts`

## Likely Primary Files
- `plugins/openapi-sdk-go/pkg/generate/refs.go`
  Evidence: query term match in path: openapi; query term match in path: openapi-sdk; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/refs.go`
  Evidence: query term match in path: openapi; query term match in path: openapi-sdk; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/client.go`
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-go/pkg/generate/client.go`
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-go/pkg/generate/types.go`
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-go/pkg/generate/generate.go` - plugins/openapi-sdk-go/pkg/generate/generate.go (go)
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-ts/pkg/generate/generate.go` - plugins/openapi-sdk-ts/pkg/generate/generate.go (go)
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk

## Likely Tests
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go`
  Evidence: query term match in path: openapi; query term match in path: openapi-sdk; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go`
  Evidence: query term match in path: openapi; query term match in path: openapi-sdk; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: openapi
- `plugins/openapi-sdk-go/pkg/generate/compile_test.go`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: openapi
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go#L14` - TestUnsupportedRef_FailsWithOneLine
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go#L14` - TestUnsupportedRef_FailsWithOneLine
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-ts/pkg/generate/compile_test.go#L21` - TestCompileTestSuite
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: openapi-sdk
- `scripts/test-recipe.sh`
  Evidence: relationship expansion: source_manifest_relationship_recovery; query term match in path: recipe; query term match in body: kalo
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: openapi-sdk

## Likely Docs / Plans / Config
None found in the initial preflight.

## Supporting Context
None found in the initial preflight.

## Related Git Receipts
- `f5ffed5` 2026-11-14 - openapi-sdk-ts/go: refuse schema $refs outside #/components/schemas with one line and a non-zero exit (experiment ope...
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/generate.go`, `plugins/openapi-sdk-go/pkg/generate/refs.go`, `plugins/openapi-sdk-go/pkg/generate/refs_test.go`
- `5d39ace` 2026-11-14 - openapi-sdk-ts/go: ApiError clash, required-before-optional params, only used Go imports, OpenAPI 3.1 warning (experi...
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/client.go`, `plugins/openapi-sdk-go/pkg/generate/generate.go`, `plugins/openapi-sdk-ts/pkg/generate/client.go`
- `4573f58` 2026-11-14 - openapi-sdk-ts/go: type aliases for non-object schemas, identifier sanitising, edge fixture + goldens
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/client.go`, `plugins/openapi-sdk-go/pkg/generate/types.go`, `plugins/openapi-sdk-ts/pkg/generate/client.go`

## Noise Risks
None found in the initial preflight.

## Freshness Warnings
These on-disk paths match the task wording but were not present in the indexed candidate set. Treat them as stale-index risk, not proof that the initial pack is wrong.

- `recipes/openapi-sdk/descriptors/namespaces/kalox.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk, kalo
- `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, plugins
- `plugins/openapi-sdk-go/plugin.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, plugins
- `plugins/openapi-sdk-ts/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, plugins
- `plugins/openapi-sdk-ts/plugin.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, plugins
- `recipes/openapi-morphe/descriptors/namespaces/kalox.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, kalo
- `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk
- `recipes/openapi-sdk/descriptors/bindings/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk

## Risk Cards
Evidence-backed checks to run before trusting the initial task context. These are not required edit targets.

- On-disk paths matched the task but were not indexed [medium, freshness]
  Agent check: Inspect the warned files or refresh the index before trusting missing context.
  Evidence: `recipes/openapi-sdk/descriptors/namespaces/kalox.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk, kalo; `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, plugins; `plugins/openapi-sdk-go/plugin.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, plugins; `plugins/openapi-sdk-ts/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, plugins

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
- found 7 likely primary file(s)
- found 10 likely test file(s)
- found 3 related Git receipt(s)

Agent instruction:
Validate the test and integration surface before editing. Record critical misses and distracting inclusions in the slice result or a task checkpoint.

## Suggested Starting Slice
Use `B01-recipe-openapi-sdk-0-1-3-a-200-201-response-that-is-an-inline-plan.md` as the first bounded plan in this task thread. Refine it before editing if primary files, tests, or integration points look incomplete.

## Agent Preflight Checklist
- [ ] Verify the likely primary files against the repo before editing.
- [ ] Search for same-package or same-command tests if test confidence is not high.
- [ ] Check receipt-touched related files before assuming the pack is complete.
- [ ] Record files actually read, edited, tests run, misses, and noise in `B01-recipe-openapi-sdk-0-1-3-a-200-201-response-that-is-an-inline-result.md` or `ds task checkpoint`.
