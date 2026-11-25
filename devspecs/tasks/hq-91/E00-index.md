# Task hq-91

## Task
One cold fetch of recipe 0.2.0 on agent.5 (clean HOME), logged the same wake; figures for the FILL commit and the recipe README

## Status
packed

## Series
E

## Profile
code-change

## Created At
2026-11-25T08:52:20Z

## Original Query
One cold fetch of recipe 0.2.0 on agent.5 (clean HOME), logged the same wake; figures for the FILL commit and the recipe README

## Repo / Workspace
- Repo: `/data/workspaces/eng1/src/kalo-plugins`
- Workspace: `/data/workspaces/eng1/src/kalo-plugins/devspecs/tasks/hq-91`

## Resources
- `task.json`
- `E01-one-cold-fetch-of-recipe-0-2-0-on-agent-5-clean-home-logged-the-plan.md`
- `E01-one-cold-fetch-of-recipe-0-2-0-on-agent-5-clean-home-logged-the-result.md`

## Task Slices
- E01: One cold fetch of recipe 0.2.0 on agent.5 (clean HOME), logged the same wake; figures for the FILL commit and the recipe README. Plan: `E01-one-cold-fetch-of-recipe-0-2-0-on-agent-5-clean-home-logged-the-plan.md`. Result: `E01-one-cold-fetch-of-recipe-0-2-0-on-agent-5-clean-home-logged-the-result.md`.

## Relevant Map Areas
- `plugins`
- `scripts`

## Likely Primary Files
- `scripts/pack-recipe.sh`
  Evidence: anchor-first ranking: score 24.000; matches recipe; fields path, title, body; query term match in path: recipe; query term match in body: fetch
- `scripts/stage-recipe.sh`
  Evidence: anchor-first ranking: score 24.000; matches recipe; fields path, title, body; query term match in path: recipe; query term match in body: fill
- `plugins/openapi-sdk-ts/pkg/generate/client.go` - plugins/openapi-sdk-ts/pkg/generate/client.go (go)
  Evidence: query term match in body: fetch; query term match in body: one; query term match in body: same
- `plugins/openapi-sdk-ts/cmd/plugin/main.go` - plugins/openapi-sdk-ts/cmd/plugin/main.go (go)
  Evidence: relationship expansion: source_manifest_family_recovery
- `plugins/openapi-sdk-go/cmd/plugin/main.go` - plugins/openapi-sdk-go/cmd/plugin/main.go (go)
  Evidence: relationship expansion: source_manifest_family_recovery
- `plugins/openapi-sdk/cmd/plugin/main.go` - plugins/openapi-sdk/cmd/plugin/main.go (go)
  Evidence: query term match in body: one; query term match in body: recipe; query term match in body: same
- `plugins/openapi-sdk-go/pkg/generate/client.go` - plugins/openapi-sdk-go/pkg/generate/client.go (go)
  Evidence: relationship expansion: source_manifest_family_recovery

## Likely Tests
- `scripts/test-recipe.sh`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: recipe
- `plugins/openapi-sdk/cmd/plugin/main_test.go#L50` - TestTargetsMatchRecipeGolden
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in title: recipe
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go#L14` - TestUnsupportedRef_FailsWithOneLine
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in title: one
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go#L14` - TestUnsupportedRef_FailsWithOneLine
  Evidence: relationship expansion: source_manifest_relationship_recovery; query term match in title: one

## Likely Docs / Plans / Config
None found in the initial preflight.

## Supporting Context
None found in the initial preflight.

## Related Git Receipts
- `b6ea57a` 2026-11-18 - openapi-sdk-ts/go: inline array of $ref typed, data field unwrapped; recipe 0.1.3 (hq-68) (#17)
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/client.go`, `plugins/openapi-sdk-ts/pkg/generate/client.go`
- `ad9bcf3` 2026-11-25 - openapi-sdk recipe 0.2.0: one wasm for both generators (hq-78) (#20)
  Matched paths: `plugins/openapi-sdk/cmd/plugin/main.go`
- `d99d1b4` 2026-11-14 - chore: monorepo layout; move plugin-openapi-sdk-ts in as plugins/openapi-sdk-ts (#1)
  Matched paths: `plugins/openapi-sdk-ts/cmd/plugin/main.go`, `plugins/openapi-sdk-ts/pkg/generate/client.go`

## Noise Risks
- `recipes/openapi-sdk/testdata/golden/ts/client.ts` - recipes/openapi-sdk/testdata/golden/ts/client.ts (typescript)
  Evidence: query term match in path: recipe; query term match in body: fetch; query term match in body: one

## Freshness Warnings
These on-disk paths match the task wording but were not present in the indexed candidate set. Treat them as stale-index risk, not proof that the initial pack is wrong.

- `recipes/openapi-morphe/README.md` - on-disk path matched task terms but was not in the indexed candidate set: recipe, readme
- `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: recipe, readme

## Risk Cards
Evidence-backed checks to run before trusting the initial task context. These are not required edit targets.

- On-disk paths matched the task but were not indexed [medium, freshness]
  Agent check: Inspect the warned files or refresh the index before trusting missing context.
  Evidence: `recipes/openapi-morphe/README.md` - on-disk path matched task terms but was not in the indexed candidate set: recipe, readme; `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: recipe, readme

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
- Noise risk: medium
- Pack completeness: medium

Why:
- found 7 likely primary file(s)
- found 4 likely test file(s)
- found 3 related Git receipt(s)
- 1 file(s) were downgraded as likely noise

Agent instruction:
Validate the test and integration surface before editing. Record critical misses and distracting inclusions in the slice result or a task checkpoint.

## Suggested Starting Slice
Use `E01-one-cold-fetch-of-recipe-0-2-0-on-agent-5-clean-home-logged-the-plan.md` as the first bounded plan in this task thread. Refine it before editing if primary files, tests, or integration points look incomplete.

## Agent Preflight Checklist
- [ ] Verify the likely primary files against the repo before editing.
- [ ] Search for same-package or same-command tests if test confidence is not high.
- [ ] Check receipt-touched related files before assuming the pack is complete.
- [ ] Record files actually read, edited, tests run, misses, and noise in `E01-one-cold-fetch-of-recipe-0-2-0-on-agent-5-clean-home-logged-the-result.md` or `ds task checkpoint`.
