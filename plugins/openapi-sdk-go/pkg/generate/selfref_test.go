package generate_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-go/pkg/generate"
	"github.com/stretchr/testify/require"
)

func selfRefTypes(t *testing.T, name string) (string, error) {
	t.Helper()
	_, filename, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(filepath.Join(filepath.Dir(filename), "..", "..", "testdata", "selfref", name))
	require.NoError(t, err)
	doc, err := generate.ParseSpec(b)
	if err != nil {
		return "", err
	}
	_, err = generate.GenerateClient(doc, "sdk")
	require.NoError(t, err)
	types, err := generate.GenerateTypes(doc, "sdk")
	require.NoError(t, err)
	return types, nil
}

// An optional self-reference (a tree's parent / children) generates instead of
// recursing until the plugin runs out of stack or memory (recipe 0.1.3 did).
func TestSelfRef_OptionalDirect(t *testing.T) {
	types, err := selfRefTypes(t, "self-ref-optional.json")
	require.NoError(t, err)
	require.Regexp(t, `Parent\s+\*Category\s+`+"`"+`json:"parent,omitempty"`, types)
	require.Regexp(t, `Children\s+\[\]Category\s+`+"`"+`json:"children,omitempty"`, types)
}

func TestSelfRef_OptionalIndirect(t *testing.T) {
	types, err := selfRefTypes(t, "indirect-optional-cycle.json")
	require.NoError(t, err)
	require.Regexp(t, `Comments\s+\[\]Comment\s+`, types)
	require.Regexp(t, `Post\s+\*Post\s+`, types)
}

// A required cycle has no finite value; the parser's circular-reference error stays.
func TestSelfRef_RequiredCycleStillFails(t *testing.T) {
	_, err := selfRefTypes(t, "required-cycle.json")
	require.ErrorContains(t, err, "circular reference")
}

// allOf is not read by the generator: the property stays interface{} (unchanged from 0.1.3).
func TestSelfRef_AllOf(t *testing.T) {
	types, err := selfRefTypes(t, "allof-self-ref.json")
	require.NoError(t, err)
	require.Regexp(t, `Manager\s+\*interface\{\}\s+`, types)
}

// The generated package for an optional self-reference passes go vet, and a nested
// value (a category with a parent and a child) survives a JSON round trip.
func TestSelfRef_GeneratedPackageBuildsAndRoundTrips(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	td := filepath.Join(filepath.Dir(filename), "..", "..", "testdata", "selfref")
	for _, fixture := range []string{"self-ref-optional.json", "indirect-optional-cycle.json"} {
		in, out := t.TempDir(), t.TempDir()
		b, err := os.ReadFile(filepath.Join(td, fixture))
		require.NoError(t, err)
		require.NoError(t, os.WriteFile(filepath.Join(in, "openapi.json"), b, 0o644))
		cfg := generate.Config{InputDir: in, OutputDir: out, SpecFileName: "openapi.json", PackageName: "sdk", GenerateTypes: true, GenerateClient: true}
		cfg.Resolve()
		require.NoError(t, generate.GenerateSDK(cfg), fixture)
		require.NoError(t, os.WriteFile(filepath.Join(out, "go.mod"), []byte("module example.com/sdk\n\ngo 1.21\n"), 0o644))
		if fixture == "self-ref-optional.json" {
			require.NoError(t, os.WriteFile(filepath.Join(out, "roundtrip_test.go"), []byte(selfRefRoundTrip), 0o644))
		}
		for _, args := range [][]string{{"vet", "./..."}, {"test", "./..."}} {
			cmd := exec.Command("go", args...)
			cmd.Dir = out
			o, err := cmd.CombinedOutput()
			require.NoError(t, err, "%s: go %v\n%s", fixture, args, o)
		}
	}
}

const selfRefRoundTrip = `package sdk

import (
	"encoding/json"
	"testing"
)

func TestNested(t *testing.T) {
	in := ` + "`" + `{"id":"b","name":"Books","parent":{"id":"r","name":"Root"},"children":[{"id":"s","name":"Sci-fi","children":[]}]}` + "`" + `
	var c Category
	if err := json.Unmarshal([]byte(in), &c); err != nil {
		t.Fatal(err)
	}
	if c.Parent == nil || *c.Parent.Name != "Root" || len(c.Children) != 1 || *c.Children[0].Id != "s" {
		t.Fatalf("got %+v", c)
	}
	out, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	var back Category
	if err := json.Unmarshal(out, &back); err != nil || *back.Parent.Id != "r" {
		t.Fatalf("round trip: %s %v", out, err)
	}
}
`
