# Task hq-61

## Task
openapi-morphe recipe in kalo-plugins: the command under the morphe try-it panel (ruling d2 of task 60)

## Status
packed

## Series
A

## Profile
code-change

## Created At
2026-11-17T15:58:26Z

## Original Query
openapi-morphe recipe in kalo-plugins: the command under the morphe try-it panel (ruling d2 of task 60)

## Repo / Workspace
- Repo: `/data/workspaces/eng_head/kalo-plugins`
- Workspace: `/data/workspaces/eng_head/kalo-plugins/devspecs/tasks/hq-61`

## Resources
- `task.json`
- `A01-openapi-morphe-recipe-in-kalo-plugins-the-command-under-the-plan.md`
- `A01-openapi-morphe-recipe-in-kalo-plugins-the-command-under-the-result.md`

## Task Slices
- A01: openapi-morphe recipe in kalo-plugins: the command under the morphe try-it panel (ruling d2 of task 60). Plan: `A01-openapi-morphe-recipe-in-kalo-plugins-the-command-under-the-plan.md`. Result: `A01-openapi-morphe-recipe-in-kalo-plugins-the-command-under-the-result.md`.

## Relevant Map Areas
- `plugins`
- `scripts`

## Likely Primary Files
- `scripts/pack-recipe.sh`
  Evidence: anchor-first ranking: score 24.000; matches recipe; fields path, title, body; query term match in path: recipe; query term match in body: kalo
- `scripts/stage-recipe.sh`
  Evidence: anchor-first ranking: score 24.000; matches recipe; fields path, title, body; query term match in path: recipe; query term match in body: openapi
- `plugins/openapi-sdk-go/pkg/generate/generate.go` - plugins/openapi-sdk-go/pkg/generate/generate.go (go)
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: try
- `plugins/openapi-sdk-ts/pkg/generate/generate.go` - plugins/openapi-sdk-ts/pkg/generate/generate.go (go)
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/refs.go` - plugins/openapi-sdk-go/pkg/generate/refs.go (go)
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/refs.go` - plugins/openapi-sdk-ts/pkg/generate/refs.go (go)
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/warn.go` - plugins/openapi-sdk-go/pkg/generate/warn.go (go)
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/warn.go` - plugins/openapi-sdk-ts/pkg/generate/warn.go (go)
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins

## Likely Tests
- `plugins/openapi-sdk-ts/pkg/generate/edge_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: command
- `plugins/openapi-sdk-ts/pkg/generate/shapes_test.go`
  Evidence: query term match in path: openapi; query term match in path: plugins; query term match in body: command
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: openapi
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/shapes_test.go`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: openapi
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `scripts/test-recipe.sh`
  Evidence: relationship expansion: source_manifest_relationship_recovery; query term match in path: recipe; query term match in body: kalo
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go#L14` - TestUnsupportedRef_FailsWithOneLine
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go#L14` - TestUnsupportedRef_FailsWithOneLine
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: openapi; query term match in path: plugins

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
  Matched paths: `plugins/openapi-sdk-ts/pkg/generate/generate.go`

## Noise Risks
None found in the initial preflight.

## Freshness Warnings
These on-disk paths match the task wording but were not present in the indexed candidate set. Treat them as stale-index risk, not proof that the initial pack is wrong.

- `recipes/openapi-sdk/descriptors/namespaces/kalox.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, recipe, kalo
- `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, plugins
- `plugins/openapi-sdk-go/plugin.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, plugins
- `plugins/openapi-sdk-ts/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, plugins
- `plugins/openapi-sdk-ts/plugin.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, plugins
- `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, recipe
- `recipes/openapi-sdk/descriptors/bindings/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, recipe
- `recipes/openapi-sdk/descriptors/bindings/openapi.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, recipe

## Risk Cards
Evidence-backed checks to run before trusting the initial task context. These are not required edit targets.

- On-disk paths matched the task but were not indexed [medium, freshness]
  Agent check: Inspect the warned files or refresh the index before trusting missing context.
  Evidence: `recipes/openapi-sdk/descriptors/namespaces/kalox.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, recipe, kalo; `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, plugins; `plugins/openapi-sdk-go/plugin.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, plugins; `plugins/openapi-sdk-ts/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, plugins

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
- found 8 likely primary file(s)
- found 9 likely test file(s)
- found 3 related Git receipt(s)

Agent instruction:
Validate the test and integration surface before editing. Record critical misses and distracting inclusions in the slice result or a task checkpoint.

## Suggested Starting Slice
Use `A01-openapi-morphe-recipe-in-kalo-plugins-the-command-under-the-plan.md` as the first bounded plan in this task thread. Refine it before editing if primary files, tests, or integration points look incomplete.

## Agent Preflight Checklist
- [ ] Verify the likely primary files against the repo before editing.
- [ ] Search for same-package or same-command tests if test confidence is not high.
- [ ] Check receipt-touched related files before assuming the pack is complete.
- [ ] Record files actually read, edited, tests run, misses, and noise in `A01-openapi-morphe-recipe-in-kalo-plugins-the-command-under-the-result.md` or `ds task checkpoint`.
