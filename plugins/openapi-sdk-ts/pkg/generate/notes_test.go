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

// TestNotesSpec_Golden pins three shapes that broke generated code: a schema named ApiError
// next to the client error class, an optional query parameter next to a required body
// (TypeScript), and no operation that marshals or decodes JSON (Go: unused imports).
// Regenerate the golden files with UPDATE_GOLDEN=1.
func TestNotesSpec_Golden(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	td := filepath.Join(filepath.Dir(filename), "..", "..", "testdata")
	out := t.TempDir()
	cfg := generate.Config{
		InputDir:       filepath.Join(td, "notes-input"),
		OutputDir:      out,
		SpecFileName:   "openapi.json",
		GenerateTypes:  true,
		GenerateClient: true,
	}
	cfg.Resolve()
	require.NoError(t, generate.GenerateSDK(cfg))

	for _, name := range []string{"types.ts", "client.ts"} {
		got, err := os.ReadFile(filepath.Join(out, name))
		require.NoError(t, err)
		gt := filepath.Join(td, "notes-ground-truth", name)
		if os.Getenv("UPDATE_GOLDEN") != "" {
			require.NoError(t, os.WriteFile(gt, got, 0o644))
		}
		want, err := os.ReadFile(gt)
		require.NoError(t, err)
		require.Equal(t, string(want), string(got), name)
	}

	// If tsc is installed, the output must also compile under --strict.
	if tsc, err := exec.LookPath("tsc"); err == nil {
		cmd := exec.Command(tsc, "--strict", "--noEmit", "--target", "es2020", "--lib", "es2020,dom",
			"--moduleResolution", "node", "types.ts", "client.ts", "index.ts")
		cmd.Dir = out
		o, err := cmd.CombinedOutput()
		require.NoError(t, err, string(o))

		// Runtime: the client returns the parsed list for an inline array of $ref, and the
		// data field (not the wrapper) for {data: [$ref]} and {data: $ref}.
		cmd = exec.Command(tsc, "--strict", "--target", "es2020", "--lib", "es2020,dom", "--module", "commonjs",
			"--moduleResolution", "node", "--outDir", "js", "types.ts", "client.ts", "index.ts")
		cmd.Dir = out
		o, err = cmd.CombinedOutput()
		require.NoError(t, err, string(o))
		require.NoError(t, os.WriteFile(filepath.Join(out, "js", "runtime.js"), []byte(notesRuntimeJS), 0o644))
		cmd = exec.Command("node", "runtime.js")
		cmd.Dir = filepath.Join(out, "js")
		o, err = cmd.CombinedOutput()
		require.NoError(t, err, string(o))
		require.Contains(t, string(o), "runtime ok")
	} else if os.Getenv("CI") != "" {
		t.Fatal("tsc not on PATH in CI: scripts/ci.sh installs it")
	} else {
		t.Log("tsc not on PATH: skipped compile check")
	}
}

const notesRuntimeJS = `const assert = require("assert");
const { Client } = require("./client");
const bodies = {
  "/notes": [{ id: "n1", title: "a" }, { id: "n2", title: "b" }],
  "/notes/page": { data: [{ id: "n3", title: "c" }], next: "x" },
  "/notes/latest": { data: { id: "n4", title: "d" } },
};
const fakeFetch = async (url) => {
  const path = new URL(url).pathname;
  return { ok: true, status: 200, json: async () => bodies[path], text: async () => "" };
};
(async () => {
  const c = new Client("http://x", fakeFetch);
  assert.deepStrictEqual((await c.listNotes()).map((n) => n.id), ["n1", "n2"]);
  assert.deepStrictEqual((await c.listNotesPage()).map((n) => n.id), ["n3"]);
  assert.strictEqual((await c.getLatestNote()).id, "n4");
  console.log("runtime ok");
})().catch((e) => { console.error(e); process.exit(1); });
`
