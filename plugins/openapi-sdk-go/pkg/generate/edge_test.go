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

	for _, name := range []string{"types.go", "client.go"} {
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

	// The output must also build as a standalone package (the go.mod is test-only).
	require.NoError(t, os.WriteFile(filepath.Join(out, "go.mod"), []byte("module example.com/sdk\n\ngo 1.21\n"), 0o644))
	cmd := exec.Command("go", "vet", "./...")
	cmd.Dir = out
	o, err := cmd.CombinedOutput()
	require.NoError(t, err, string(o))
}
