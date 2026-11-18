package generate_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-ts/pkg/generate"
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
	_ = generate.GenerateClient(doc, "./types")
	return generate.GenerateTypes(doc), nil
}

// An optional self-reference (a tree's parent / children) generates instead of
// recursing until the plugin runs out of stack or memory (recipe 0.1.3 did).
func TestSelfRef_OptionalDirect(t *testing.T) {
	types, err := selfRefTypes(t, "self-ref-optional.json")
	require.NoError(t, err)
	require.Contains(t, types, "  parent?: Category;\n")
	require.Contains(t, types, "  children?: Category[];\n")
}

func TestSelfRef_OptionalIndirect(t *testing.T) {
	types, err := selfRefTypes(t, "indirect-optional-cycle.json")
	require.NoError(t, err)
	require.Contains(t, types, "  comments?: Comment[];\n")
	require.Contains(t, types, "  post?: Post;\n")
}

// A required cycle has no finite value; the parser's circular-reference error stays.
func TestSelfRef_RequiredCycleStillFails(t *testing.T) {
	_, err := selfRefTypes(t, "required-cycle.json")
	require.ErrorContains(t, err, "circular reference")
}

// allOf is not read by the generator: the property stays unknown (unchanged from 0.1.3).
func TestSelfRef_AllOf(t *testing.T) {
	types, err := selfRefTypes(t, "allof-self-ref.json")
	require.NoError(t, err)
	require.Contains(t, types, "  manager?: unknown;\n")
}

// The generated SDK for an optional self-reference compiles under --strict, and a
// nested category (parent and child) comes back through the client from a stub fetch.
func TestSelfRef_GeneratedSDKCompilesAndRuns(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(filepath.Join(filepath.Dir(filename), "..", "..", "testdata", "selfref", "self-ref-optional.json"))
	require.NoError(t, err)
	in, out := t.TempDir(), t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(in, "openapi.json"), b, 0o644))
	cfg := generate.Config{InputDir: in, OutputDir: out, SpecFileName: "openapi.json", GenerateTypes: true, GenerateClient: true}
	cfg.Resolve()
	require.NoError(t, generate.GenerateSDK(cfg))

	tsc, err := exec.LookPath("tsc")
	if err != nil {
		if os.Getenv("CI") != "" {
			t.Fatal("tsc not on PATH in CI: scripts/ci.sh installs it")
		}
		t.Skip("tsc not on PATH: skipped compile check")
	}
	cmd := exec.Command(tsc, "--strict", "--target", "es2020", "--lib", "es2020,dom", "--module", "commonjs",
		"--moduleResolution", "node", "--outDir", "js", "types.ts", "client.ts", "index.ts")
	cmd.Dir = out
	o, err := cmd.CombinedOutput()
	require.NoError(t, err, string(o))
	require.NoError(t, os.WriteFile(filepath.Join(out, "js", "runtime.js"), []byte(selfRefRuntimeJS), 0o644))
	cmd = exec.Command("node", "runtime.js")
	cmd.Dir = filepath.Join(out, "js")
	o, err = cmd.CombinedOutput()
	require.NoError(t, err, string(o))
	require.Contains(t, string(o), "runtime ok")
}

const selfRefRuntimeJS = `const assert = require("assert");
const { Client } = require("./client");
const body = { id: "b", name: "Books", parent: { id: "r", name: "Root" }, children: [{ id: "s", name: "Sci-fi", children: [] }] };
const fakeFetch = async () => ({ ok: true, status: 200, json: async () => body, text: async () => "" });
(async () => {
  const c = await new Client("http://x", fakeFetch).getCategory("b");
  assert.strictEqual(c.parent.name, "Root");
  assert.strictEqual(c.children[0].id, "s");
  console.log("runtime ok");
})().catch((e) => { console.error(e); process.exit(1); });
`
