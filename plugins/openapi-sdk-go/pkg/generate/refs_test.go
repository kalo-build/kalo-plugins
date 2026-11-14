package generate_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-go/pkg/generate"
	"github.com/stretchr/testify/require"
)

// A schema $ref outside #/components/schemas must stop the run with one line and write nothing.
func TestUnsupportedRef_FailsWithOneLine(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	td := filepath.Join(filepath.Dir(filename), "..", "..", "testdata")
	out := filepath.Join(t.TempDir(), "out")
	cfg := generate.Config{
		InputDir:       filepath.Join(td, "badref-input"),
		OutputDir:      out,
		SpecFileName:   "openapi.yaml",
		GenerateTypes:  true,
		GenerateClient: true,
	}
	cfg.Resolve()
	err := generate.GenerateSDK(cfg)
	require.Error(t, err)
	require.Equal(t, "unsupported $ref \"#/paths/~1pets/get/responses/200/content/application~1json/schema/items\" (3 such $refs in the schemas read): only $refs to #/components/schemas/<Name> are supported", err.Error())
	require.NotContains(t, err.Error(), "\n")
	_, statErr := os.Stat(out)
	require.True(t, os.IsNotExist(statErr), "output directory must not be created")
}
