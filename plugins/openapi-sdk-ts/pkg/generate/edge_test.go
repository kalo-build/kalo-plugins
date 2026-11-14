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

// TestEdgeSpec_Golden covers what real specs contain and the hello-world fixture does not:
// non-object component schemas (array alias, enum), schema and operation names with
// punctuation, reserved-word and digit-leading parameters, missing operationIds that
// would collide. Regenerate the golden files with UPDATE_GOLDEN=1.
func TestEdgeSpec_Golden(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	td := filepath.Join(filepath.Dir(filename), "..", "..", "testdata")
	out := t.TempDir()
	cfg := generate.Config{
		InputDir:       filepath.Join(td, "edge-input"),
		OutputDir:      out,
		SpecFileName:   "openapi.yaml",
		GenerateTypes:  true,
		GenerateClient: true,
	}
	cfg.Resolve()
	require.NoError(t, generate.GenerateSDK(cfg))

	for _, name := range []string{"types.ts", "client.ts"} {
		got, err := os.ReadFile(filepath.Join(out, name))
		require.NoError(t, err)
		gt := filepath.Join(td, "edge-ground-truth", name)
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
	} else {
		t.Log("tsc not on PATH: skipped compile check")
	}
}
