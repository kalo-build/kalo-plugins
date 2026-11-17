# Task hq-59 B01 Plan

## Goal
Prepare, don't merge: kalo-cli README (reach §2c) and page links in the kalo-plugins READMEs, as held PRs for "A: merged"

## Description
Create a bounded implementation slice for `Prepare, don't merge: kalo-cli README (reach §2c) and page links in the kalo-plugins READMEs, as held PRs for "A: merged"`. This plan is grounded by the task index preflight, but it is not authoritative; confirm predicted files and tests before making edits.

## Resources
- `B00-index.md`
- `B01-prepare-don-t-merge-kalo-cli-readme-reach-2c-and-page-links-in-result.md`
- `task.json`
- `plugins/openapi-sdk-go/pkg/generate/client.go`
- `plugins/openapi-sdk-ts/pkg/generate/client.go`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `scripts/test-recipe.sh`
- `plugins/openapi-sdk-go/pkg/generate/shapes_test.go`
- `plugins/openapi-sdk-go/pkg/generate/compile_test.go`
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`

## Starting Context
### Files to Inspect First
- `plugins/openapi-sdk-go/pkg/generate/client.go`
- `plugins/openapi-sdk-ts/pkg/generate/client.go`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`

### Tests to Inspect First
- `scripts/test-recipe.sh`
- `plugins/openapi-sdk-go/pkg/generate/shapes_test.go`
- `plugins/openapi-sdk-go/pkg/generate/compile_test.go`
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L18`

## Expected Change Surface
- `plugins/openapi-sdk-go/pkg/generate/client.go`
- `plugins/openapi-sdk-ts/pkg/generate/client.go`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`

## Out-of-Scope Areas
- Replanning the whole thread unless evidence says this slice should split or be superseded.
- Broad pack-ranking changes unless they are necessary for this task.
- Treating the generated context as complete without verification.

## Risks
- Task-related on-disk paths may be missing from the indexed candidate set.
- Pack completeness is not high; verify the working set before editing.
- On-disk paths matched the task but were not indexed: Inspect the warned files or refresh the index before trusting missing context. Evidence: `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: readme, plugins; `plugins/openapi-sdk-ts/README.md` - on-disk path matched task terms but was not in the indexed candidate set: readme, plugins.
- Initial pack includes downgraded noise candidates; avoid editing them unless verification supports it.

## Success Criteria
- [ ] Primary implementation surface is verified before edits.
- [ ] Relevant tests are found or the test-surface miss is recorded.
- [ ] Changes stay inside the bounded slice.
- [ ] A checkpoint records actual files, tests, misses, noise, and decision.

## Tasks
- [ ] Inspect the predicted primary files.
- [ ] Inspect same-package, same-stem, or receipt-related tests.
- [ ] Refine the slice if context is incomplete.
- [ ] Implement the smallest useful change.
- [ ] Run focused validation.
- [ ] Update `B01-prepare-don-t-merge-kalo-cli-readme-reach-2c-and-page-links-in-result.md` or run `ds task checkpoint`.

## Decision Gates
- Promote: the workspace was useful enough and misses are actionable.
- Improve: useful start, but incomplete/noisy enough to require template or retrieval changes.
- Rework: task workspace feels like planning overhead or fails to capture useful evidence.
- Rollback: workspace creates false confidence or worsens agent performance.
- Block: external input or a missing prerequisite prevents useful progress.
