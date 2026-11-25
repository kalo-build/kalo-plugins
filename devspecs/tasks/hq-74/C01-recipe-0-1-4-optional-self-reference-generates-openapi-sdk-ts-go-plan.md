# Task hq-74 C01 Plan

## Goal
Recipe 0.1.4: optional self-reference generates (openapi-sdk-ts/go), PR only, no tag before the website merge

## Description
Create a bounded implementation slice for `Recipe 0.1.4: optional self-reference generates (openapi-sdk-ts/go), PR only, no tag before the website merge`. This plan is grounded by the task index preflight, but it is not authoritative; confirm predicted files and tests before making edits.

## Resources
- `C00-index.md`
- `C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-result.md`
- `task.json`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-ts/pkg/generate/warn.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L57`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L80`
- `plugins/openapi-sdk-ts/pkg/generate/notes_test.go`

## Starting Context
### Files to Inspect First
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-ts/pkg/generate/warn.go`

### Tests to Inspect First
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L57`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L80`
- `plugins/openapi-sdk-ts/pkg/generate/notes_test.go`
- `plugins/openapi-sdk-ts/pkg/generate/shapes_test.go`
- `plugins/openapi-sdk-ts/pkg/generate/compile_test.go#L21`
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18`
- `scripts/test-recipe.sh`
- `plugins/openapi-sdk-ts/pkg/generate/warn_test.go#L8`

## Expected Change Surface
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-ts/pkg/generate/warn.go`

## Out-of-Scope Areas
- Replanning the whole thread unless evidence says this slice should split or be superseded.
- Broad pack-ranking changes unless they are necessary for this task.
- Treating the generated context as complete without verification.

## Risks
- Task-related on-disk paths may be missing from the indexed candidate set.
- Pack completeness is not high; verify the working set before editing.
- On-disk paths matched the task but were not indexed: Inspect the warned files or refresh the index before trusting missing context. Evidence: `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk; `recipes/openapi-sdk/descriptors/bindings/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: recipe, openapi, sdk.

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
- [ ] Update `C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-result.md` or run `ds task checkpoint`.

## Decision Gates
- Promote: the workspace was useful enough and misses are actionable.
- Improve: useful start, but incomplete/noisy enough to require template or retrieval changes.
- Rework: task workspace feels like planning overhead or fails to capture useful evidence.
- Rollback: workspace creates false confidence or worsens agent performance.
- Block: external input or a missing prerequisite prevents useful progress.
