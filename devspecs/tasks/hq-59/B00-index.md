# Task hq-59

## Task
Prepare, don't merge: kalo-cli README (reach §2c) and page links in the kalo-plugins READMEs, as held PRs for "A: merged"

## Status
packed

## Series
B

## Profile
code-change

## Created At
2026-11-17T15:42:03Z

## Original Query
Prepare, don't merge: kalo-cli README (reach §2c) and page links in the kalo-plugins READMEs, as held PRs for "A: merged"

## Repo / Workspace
- Repo: `/tmp/kp`
- Workspace: `/tmp/kp/devspecs/tasks/hq-59`

## Resources
- `task.json`
- `B01-prepare-don-t-merge-kalo-cli-readme-reach-2c-and-page-links-in-plan.md`
- `B01-prepare-don-t-merge-kalo-cli-readme-reach-2c-and-page-links-in-result.md`

## Task Slices
- B01: Prepare, don't merge: kalo-cli README (reach §2c) and page links in the kalo-plugins READMEs, as held PRs for "A: merged". Plan: `B01-prepare-don-t-merge-kalo-cli-readme-reach-2c-and-page-links-in-plan.md`. Result: `B01-prepare-don-t-merge-kalo-cli-readme-reach-2c-and-page-links-in-result.md`.

## Relevant Map Areas
- `plugins`
- `scripts`

## Likely Primary Files
- `plugins/openapi-sdk-go/pkg/generate/client.go`
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: cli; query term match in path: plugins
- `plugins/openapi-sdk-ts/pkg/generate/client.go`
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: cli; query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/generate.go` - plugins/openapi-sdk-go/pkg/generate/generate.go (go)
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: plugins; query term match in body: cli
- `plugins/openapi-sdk-ts/pkg/generate/generate.go` - plugins/openapi-sdk-ts/pkg/generate/generate.go (go)
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: plugins; query term match in body: cli

## Likely Tests
- `scripts/test-recipe.sh`
  Evidence: query term match in body: cli; query term match in body: kalo; identifier/body match: kalo-cli
- `plugins/openapi-sdk-go/pkg/generate/shapes_test.go`
  Evidence: query term match in path: plugins; query term match in body: cli; query term match in body: kalo
- `plugins/openapi-sdk-go/pkg/generate/compile_test.go`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`
  Evidence: relationship expansion: source_manifest_test_reservation; pack tier: related (reserved manifest test with direct query evidence); query term match in path: plugins
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: plugins; query term match in body: cli
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L18` - TestGenerateSDK_Full
  Evidence: relationship expansion: source_manifest_loss_safe_preserved; query term match in path: plugins; query term match in body: cli

## Likely Docs / Plans / Config
None found in the initial preflight.

## Supporting Context
None found in the initial preflight.

## Related Git Receipts
- `5d39ace` 2026-11-14 - openapi-sdk-ts/go: ApiError clash, required-before-optional params, only used Go imports, OpenAPI 3.1 warning (experi...
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/client.go`, `plugins/openapi-sdk-go/pkg/generate/generate.go`, `plugins/openapi-sdk-go/pkg/generate/shapes_test.go`
- `4b674de` 2026-11-14 - chore: move plugin-openapi-sdk-go in as plugins/openapi-sdk-go
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/client.go`, `plugins/openapi-sdk-go/pkg/generate/compile_test.go`, `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `4573f58` 2026-11-14 - openapi-sdk-ts/go: type aliases for non-object schemas, identifier sanitising, edge fixture + goldens
  Matched paths: `plugins/openapi-sdk-go/pkg/generate/client.go`, `plugins/openapi-sdk-go/pkg/generate/edge_test.go`, `plugins/openapi-sdk-ts/pkg/generate/client.go`

## Noise Risks
- `plugins/openapi-sdk-ts/testdata/edge-ground-truth/client.ts`
  Evidence: query term match in path: cli; query term match in path: plugins
- `plugins/openapi-sdk-ts/testdata/ground-truth/client.ts` - plugins/openapi-sdk-ts/testdata/ground-truth/client.ts (typescript)
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: cli; query term match in path: plugins
- `plugins/openapi-sdk-ts/testdata/shapes-ground-truth/client.ts` - plugins/openapi-sdk-ts/testdata/shapes-ground-truth/client.ts (typescript)
  Evidence: relationship expansion: source_manifest_family_recovery; query term match in path: cli; query term match in path: plugins

## Freshness Warnings
These on-disk paths match the task wording but were not present in the indexed candidate set. Treat them as stale-index risk, not proof that the initial pack is wrong.

- `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: readme, plugins
- `plugins/openapi-sdk-ts/README.md` - on-disk path matched task terms but was not in the indexed candidate set: readme, plugins

## Risk Cards
Evidence-backed checks to run before trusting the initial task context. These are not required edit targets.

- On-disk paths matched the task but were not indexed [medium, freshness]
  Agent check: Inspect the warned files or refresh the index before trusting missing context.
  Evidence: `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: readme, plugins; `plugins/openapi-sdk-ts/README.md` - on-disk path matched task terms but was not in the indexed candidate set: readme, plugins

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
- found 4 likely primary file(s)
- found 6 likely test file(s)
- found 3 related Git receipt(s)
- 3 file(s) were downgraded as likely noise

Agent instruction:
Validate the test and integration surface before editing. Record critical misses and distracting inclusions in the slice result or a task checkpoint.

## Suggested Starting Slice
Use `B01-prepare-don-t-merge-kalo-cli-readme-reach-2c-and-page-links-in-plan.md` as the first bounded plan in this task thread. Refine it before editing if primary files, tests, or integration points look incomplete.

## Agent Preflight Checklist
- [ ] Verify the likely primary files against the repo before editing.
- [ ] Search for same-package or same-command tests if test confidence is not high.
- [ ] Check receipt-touched related files before assuming the pack is complete.
- [ ] Record files actually read, edited, tests run, misses, and noise in `B01-prepare-don-t-merge-kalo-cli-readme-reach-2c-and-page-links-in-result.md` or `ds task checkpoint`.
