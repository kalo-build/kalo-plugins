# Task hq-78 D01 Plan

## Goal
One-wasm openapi-sdk recipe (0.2.0): build to LGTM against bars; no merge, no tag

## Description
Create a bounded implementation slice for `One-wasm openapi-sdk recipe (0.2.0): build to LGTM against bars; no merge, no tag`. This plan is grounded by the task index preflight, but it is not authoritative; confirm predicted files and tests before making edits.

## Resources
- `D00-index.md`
- `D01-one-wasm-openapi-sdk-recipe-0-2-0-build-to-lgtm-against-bars-no-result.md`
- `task.json`
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-ts/pkg/generate/client.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/client.go`
- `plugins/openapi-sdk-go/pkg/generate/types.go`
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`

## Starting Context
### Files to Inspect First
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-ts/pkg/generate/client.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/client.go`
- `plugins/openapi-sdk-go/pkg/generate/types.go`
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`

### Tests to Inspect First
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`
- `plugins/openapi-sdk-go/pkg/generate/linear_test.go`
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go#L14`
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go#L14`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L18`
- `scripts/test-recipe.sh`
- `plugins/openapi-sdk-go/pkg/generate/selfref_test.go#L32`
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18`

## Expected Change Surface
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-ts/pkg/generate/client.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/client.go`
- `plugins/openapi-sdk-go/pkg/generate/types.go`
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`

## Out-of-Scope Areas
- Replanning the whole thread unless evidence says this slice should split or be superseded.
- Broad pack-ranking changes unless they are necessary for this task.
- Treating the generated context as complete without verification.

## Risks
- Task-related on-disk paths may be missing from the indexed candidate set.
- Pack completeness is not high; verify the working set before editing.
- On-disk paths matched the task but were not indexed: Inspect the warned files or refresh the index before trusting missing context. Evidence: `recipes/openapi-sdk/README.md` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe; `recipes/openapi-sdk/descriptors/bindings/go-sdk.yaml` - on-disk path matched task terms but was not in the indexed candidate set: openapi, sdk, recipe.
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
- [ ] Update `D01-one-wasm-openapi-sdk-recipe-0-2-0-build-to-lgtm-against-bars-no-result.md` or run `ds task checkpoint`.

## Decision Gates
- Promote: the workspace was useful enough and misses are actionable.
- Improve: useful start, but incomplete/noisy enough to require template or retrieval changes.
- Rework: task workspace feels like planning overhead or fails to capture useful evidence.
- Rollback: workspace creates false confidence or worsens agent performance.
- Block: external input or a missing prerequisite prevents useful progress.
