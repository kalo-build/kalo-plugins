# Task hq-61 A01 Plan

## Goal
openapi-morphe recipe in kalo-plugins: the command under the morphe try-it panel (ruling d2 of task 60)

## Description
Create a bounded implementation slice for `openapi-morphe recipe in kalo-plugins: the command under the morphe try-it panel (ruling d2 of task 60)`. This plan is grounded by the task index preflight, but it is not authoritative; confirm predicted files and tests before making edits.

## Resources
- `A00-index.md`
- `A01-openapi-morphe-recipe-in-kalo-plugins-the-command-under-the-result.md`
- `task.json`
- `scripts/pack-recipe.sh`
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/refs.go`
- `plugins/openapi-sdk-ts/pkg/generate/refs.go`
- `plugins/openapi-sdk-go/pkg/generate/warn.go`
- `plugins/openapi-sdk-ts/pkg/generate/warn.go`

## Starting Context
### Files to Inspect First
- `scripts/pack-recipe.sh`
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/refs.go`
- `plugins/openapi-sdk-ts/pkg/generate/refs.go`
- `plugins/openapi-sdk-go/pkg/generate/warn.go`
- `plugins/openapi-sdk-ts/pkg/generate/warn.go`

### Tests to Inspect First
- `plugins/openapi-sdk-ts/pkg/generate/edge_test.go`
- `plugins/openapi-sdk-ts/pkg/generate/shapes_test.go`
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18`
- `plugins/openapi-sdk-go/pkg/generate/shapes_test.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L18`
- `scripts/test-recipe.sh`
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go#L14`
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go#L14`

## Expected Change Surface
- `scripts/pack-recipe.sh`
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/refs.go`
- `plugins/openapi-sdk-ts/pkg/generate/refs.go`
- `plugins/openapi-sdk-go/pkg/generate/warn.go`
- `plugins/openapi-sdk-ts/pkg/generate/warn.go`

## Out-of-Scope Areas
- Replanning the whole thread unless evidence says this slice should split or be superseded.
- Broad pack-ranking changes unless they are necessary for this task.
- Treating the generated context as complete without verification.

## Risks
- Task-related on-disk paths may be missing from the indexed candidate set.
- Pack completeness is not high; verify the working set before editing.
- On-disk paths matched the task but were not indexed: Inspect the warned files or refresh the index before trusting missing context. Evidence: `recipes/openapi-sdk/descriptors/namespaces/kalox.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, recipe, kalo; `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, plugins.

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
- [ ] Update `A01-openapi-morphe-recipe-in-kalo-plugins-the-command-under-the-result.md` or run `ds task checkpoint`.

## Decision Gates
- Promote: the workspace was useful enough and misses are actionable.
- Improve: useful start, but incomplete/noisy enough to require template or retrieval changes.
- Rework: task workspace feels like planning overhead or fails to capture useful evidence.
- Rollback: workspace creates false confidence or worsens agent performance.
- Block: external input or a missing prerequisite prevents useful progress.
