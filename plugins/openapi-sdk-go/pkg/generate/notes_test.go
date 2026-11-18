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

	for _, name := range []string{"types.go", "client.go"} {
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

	// The output must also build as a standalone package (the go.mod is test-only).
	require.NoError(t, os.WriteFile(filepath.Join(out, "go.mod"), []byte("module example.com/sdk\n\ngo 1.21\n"), 0o644))
	cmd := exec.Command("go", "vet", "./...")
	cmd.Dir = out
	o, err := cmd.CombinedOutput()
	require.NoError(t, err, string(o))

	// Runtime: the generated client returns the parsed list for an inline array of $ref,
	// and the data field (not the wrapper) for {data: [$ref]} and {data: $ref}.
	require.NoError(t, os.WriteFile(filepath.Join(out, "runtime_test.go"), []byte(notesRuntimeTest), 0o644))
	cmd = exec.Command("go", "test", "./...")
	cmd.Dir = out
	o, err = cmd.CombinedOutput()
	require.NoError(t, err, string(o))
}

const notesRuntimeTest = `package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRuntime(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/notes":
			w.Write([]byte(` + "`" + `[{"id":"n1","title":"a"},{"id":"n2","title":"b"}]` + "`" + `))
		case "/notes/page":
			w.Write([]byte(` + "`" + `{"data":[{"id":"n3","title":"c"}],"next":"x"}` + "`" + `))
		case "/notes/latest":
			w.Write([]byte(` + "`" + `{"data":{"id":"n4","title":"d"}}` + "`" + `))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	c := &Client{BaseURL: srv.URL, HTTPClient: srv.Client()}
	ctx := context.Background()

	list, err := c.ListNotes(ctx)
	if err != nil || len(list) != 2 || list[1].Id != "n2" {
		t.Fatalf("ListNotes: %v %+v", err, list)
	}
	page, err := c.ListNotesPage(ctx)
	if err != nil || len(page) != 1 || page[0].Id != "n3" {
		t.Fatalf("ListNotesPage: %v %+v", err, page)
	}
	latest, err := c.GetLatestNote(ctx)
	if err != nil || latest.Id != "n4" || latest.Title != "d" {
		t.Fatalf("GetLatestNote: %v %+v", err, latest)
	}
}
`
