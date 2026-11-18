# Task hq-74 C01 Result

## Summary
- Target: `C01` - Recipe 0.1.4: optional self-reference generates (openapi-sdk-ts/go), PR only, no tag before the website merge
- Outcome: -

## Completion Contract
- Attempted slice: `C01` - Recipe 0.1.4: optional self-reference generates (openapi-sdk-ts/go), PR only, no tag before the website merge
- Gate tested: promote, improve, rework, rollback, or block
- What changed: -
- Evidence for decision: -
- What remains: -
- Next iteration: -

## Changed Files
-

## Tests
-

## Decision
-

## Follow-up
-

## References
- `C00-index.md`
- `C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-plan.md`

## Checkpoint History

### Checkpoint
- Created At: 2026-11-18T18:44:31Z
- Stage: validated
- Decision: promote
- Source: `checkpoints/20261118-184431-validated.md`
- Structured Evidence: `checkpoints/20261118-184431-validated.json`
- What changed: -
- Evidence for decision: 16 file(s) edited; 1 test command(s)
- What remains: -
- Next iteration: promote to the next slice
- Files edited:
  - `plugins/openapi-sdk-go/pkg/generate/generate.go`
  - `plugins/openapi-sdk-ts/pkg/generate/generate.go`
  - `devspecs/tasks/hq-74/C00-index.md`
  - `devspecs/tasks/hq-74/C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-plan.md`
  - `devspecs/tasks/hq-74/C01-recipe-0-1-4-optional-self-reference-generates-openapi-sdk-ts-go-result.md`
  - `devspecs/tasks/hq-74/task.json`
  - `plugins/openapi-sdk-go/pkg/generate/selfref_test.go`
  - `plugins/openapi-sdk-go/testdata/selfref/allof-self-ref.json`
  - `plugins/openapi-sdk-go/testdata/selfref/indirect-optional-cycle.json`
  - `plugins/openapi-sdk-go/testdata/selfref/required-cycle.json`
  - `plugins/openapi-sdk-go/testdata/selfref/self-ref-optional.json`
  - `plugins/openapi-sdk-ts/pkg/generate/selfref_test.go`
  - `plugins/openapi-sdk-ts/testdata/selfref/allof-self-ref.json`
  - `plugins/openapi-sdk-ts/testdata/selfref/indirect-optional-cycle.json`
  - `plugins/openapi-sdk-ts/testdata/selfref/required-cycle.json`
  - `plugins/openapi-sdk-ts/testdata/selfref/self-ref-optional.json`
- Tests run:
  - `go vet ./... && go test ./... in plugins/openapi-sdk-ts (tsc 5.8.2 on PATH) and plugins/openapi-sdk-go; mutation: reverting ts generate.go -> stack overflow in TestSelfRef_OptionalDirect`

### Checkpoint
- Created At: 2026-11-18T19:11:48Z
- Stage: validated
- Decision: promote
- Source: `checkpoints/20261118-191148-validated.md`
- Structured Evidence: `checkpoints/20261118-191148-validated.json`
- What changed: -
- Evidence for decision: 1 test command(s)
- What remains: -
- Next iteration: promote to the next slice
- Tests run:
  - `CI=1 go vet + go test -count=1 ./... in openapi-sdk-go and openapi-sdk-ts (tsc on PATH); wasip1 builds; mutations: c5e1c0f generator -> dense n=40 killed at 60 s; refSchema without items -> wrapper test fails`
