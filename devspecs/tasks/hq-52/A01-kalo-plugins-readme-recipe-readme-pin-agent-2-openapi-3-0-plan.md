# Task hq-52 A01 Plan

## Goal
kalo-plugins README + recipe README: pin agent.2, OpenAPI 3.0 wording, drop the stale "until published" block

## Description
Create a bounded implementation slice for `kalo-plugins README + recipe README: pin agent.2, OpenAPI 3.0 wording, drop the stale "until published" block`. This plan is grounded by the task index preflight, but it is not authoritative; confirm predicted files and tests before making edits.

## Resources
- `A00-index.md`
- `A01-kalo-plugins-readme-recipe-readme-pin-agent-2-openapi-3-0-result.md`
- `task.json`
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/refs.go`
- `plugins/openapi-sdk-go/pkg/generate/spec.go`
- `plugins/openapi-sdk-go/pkg/generate/warn.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/refs.go`
- `plugins/openapi-sdk-ts/pkg/generate/spec.go`
- `plugins/openapi-sdk-ts/pkg/generate/warn.go`

## Starting Context
### Files to Inspect First
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/refs.go`
- `plugins/openapi-sdk-go/pkg/generate/spec.go`
- `plugins/openapi-sdk-go/pkg/generate/warn.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/refs.go`
- `plugins/openapi-sdk-ts/pkg/generate/spec.go`
- `plugins/openapi-sdk-ts/pkg/generate/warn.go`
- `scripts/pack-recipe.sh`
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-go/pkg/generate/config.go`
- `plugins/openapi-sdk-ts/pkg/generate/config.go`
- `plugins/openapi-sdk-go/cmd/plugin/main.go`
- `plugins/openapi-sdk-go/pkg/generate/types.go`
- `plugins/openapi-sdk-ts/cmd/plugin/main.go`
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
- `plugins/openapi-sdk-go/pkg/generate/util.go`
- `plugins/openapi-sdk-go/pkg/generate/client.go`
- `plugins/openapi-sdk-ts/pkg/generate/util.go`
- `plugins/openapi-sdk-ts/pkg/generate/client.go`

### Tests to Inspect First
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L18`
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L61`
- `plugins/openapi-sdk-go/pkg/generate/warn_test.go#L8`
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go`
- `plugins/openapi-sdk-go/pkg/generate/compile_test.go#L21`
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go#L18`
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L86`
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L99`
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L118`
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go#L137`
- `plugins/openapi-sdk-go/pkg/generate/refs_test.go#L14`
- `plugins/openapi-sdk-go/pkg/generate/shapes_test.go#L18`
- `plugins/openapi-sdk-ts/pkg/generate/compile_test.go#L21`
- `plugins/openapi-sdk-ts/pkg/generate/edge_test.go#L18`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L18`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L57`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L80`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L103`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L131`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L144`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L161`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L178`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go#L185`
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go#L14`
- `plugins/openapi-sdk-ts/pkg/generate/shapes_test.go#L18`
- `plugins/openapi-sdk-ts/pkg/generate/warn_test.go#L8`
- `scripts/test-recipe.sh`
- `plugins/openapi-sdk-go/pkg/generate/warn_test.go`
- `plugins/openapi-sdk-ts/pkg/generate/warn_test.go`
- `plugins/openapi-sdk-go/pkg/generate/edge_test.go`
- `plugins/openapi-sdk-go/pkg/generate/shapes_test.go`
- `plugins/openapi-sdk-ts/pkg/generate/edge_test.go`
- `plugins/openapi-sdk-ts/pkg/generate/refs_test.go`
- `plugins/openapi-sdk-ts/pkg/generate/shapes_test.go`
- `plugins/openapi-sdk-go/pkg/generate/compile_test.go`
- `plugins/openapi-sdk-ts/pkg/generate/compile_test.go`
- `plugins/openapi-sdk-go/pkg/generate/generate_test.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate_test.go`

## Expected Change Surface
- `plugins/openapi-sdk-go/pkg/generate/generate.go`
- `plugins/openapi-sdk-go/pkg/generate/refs.go`
- `plugins/openapi-sdk-go/pkg/generate/spec.go`
- `plugins/openapi-sdk-go/pkg/generate/warn.go`
- `plugins/openapi-sdk-ts/pkg/generate/generate.go`
- `plugins/openapi-sdk-ts/pkg/generate/refs.go`
- `plugins/openapi-sdk-ts/pkg/generate/spec.go`
- `plugins/openapi-sdk-ts/pkg/generate/warn.go`
- `scripts/pack-recipe.sh`
- `scripts/stage-recipe.sh`
- `plugins/openapi-sdk-go/pkg/generate/config.go`
- `plugins/openapi-sdk-ts/pkg/generate/config.go`
- `plugins/openapi-sdk-go/cmd/plugin/main.go`
- `plugins/openapi-sdk-go/pkg/generate/types.go`
- `plugins/openapi-sdk-ts/cmd/plugin/main.go`
- `plugins/openapi-sdk-ts/pkg/generate/types.go`
- `plugins/openapi-sdk-go/pkg/generate/util.go`
- `plugins/openapi-sdk-go/pkg/generate/client.go`
- `plugins/openapi-sdk-ts/pkg/generate/util.go`
- `plugins/openapi-sdk-ts/pkg/generate/client.go`

## Out-of-Scope Areas
- Replanning the whole thread unless evidence says this slice should split or be superseded.
- Broad pack-ranking changes unless they are necessary for this task.
- Treating the generated context as complete without verification.

## Risks
- Task-related on-disk paths may be missing from the indexed candidate set.
- Pack completeness is not high; verify the working set before editing.
- On-disk paths matched the task but were not indexed: Inspect the warned files or refresh the index before trusting missing context. Evidence: `plugins/openapi-sdk-go/README.md` - on-disk path matched task terms but was not in the indexed candidate set: plugins, readme, openapi; `plugins/openapi-sdk-ts/README.md` - on-disk path matched task terms but was not in the indexed candidate set: plugins, readme, openapi.
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
- [ ] Update `A01-kalo-plugins-readme-recipe-readme-pin-agent-2-openapi-3-0-result.md` or run `ds task checkpoint`.

## Decision Gates
- Promote: the workspace was useful enough and misses are actionable.
- Improve: useful start, but incomplete/noisy enough to require template or retrieval changes.
- Rework: task workspace feels like planning overhead or fails to capture useful evidence.
- Rollback: workspace creates false confidence or worsens agent performance.
- Block: external input or a missing prerequisite prevents useful progress.
