package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// The recipe's fixture and golden output: one binary must give the same tree
// for each target as the recipe's two-wasm golden.
const (
	fixtureDir = "../../../../recipes/openapi-sdk/testdata"
	goldenDir  = "../../../../recipes/openapi-sdk/testdata/golden"
)

func config(t *testing.T, target, in, out string) string {
	t.Helper()
	raw, err := json.Marshal(map[string]any{
		"inputPath":  in,
		"outputPath": out,
		"config":     map[string]any{"target": target, "specFileName": "openapi.yaml"},
	})
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

func readTree(t *testing.T, dir string) map[string][]byte {
	t.Helper()
	files := map[string][]byte{}
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(dir, path)
		data, err := os.ReadFile(path)
		files[rel] = data
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func TestTargetsMatchRecipeGolden(t *testing.T) {
	for _, target := range []string{"ts", "go"} {
		t.Run(target, func(t *testing.T) {
			out := t.TempDir()
			var stderr bytes.Buffer
			if code := run([]string{"plugin", config(t, target, fixtureDir, out)}, &bytes.Buffer{}, &stderr); code != 0 {
				t.Fatalf("exit %d: %s", code, stderr.String())
			}
			want := readTree(t, filepath.Join(goldenDir, target))
			got := readTree(t, out)
			if len(got) != len(want) {
				t.Fatalf("got %d files, golden has %d", len(got), len(want))
			}
			for name, data := range want {
				if !bytes.Equal(got[name], data) {
					t.Errorf("%s/%s differs from the golden output", target, name)
				}
			}
		})
	}
}

func TestTargetRequired(t *testing.T) {
	for _, target := range []string{"", "python", "TS"} {
		var stderr bytes.Buffer
		code := run([]string{"plugin", config(t, target, fixtureDir, t.TempDir())}, &bytes.Buffer{}, &stderr)
		if code != ErrInvalidConfig || !strings.Contains(stderr.String(), "config.target") {
			t.Errorf("target %q: exit %d, stderr %q", target, code, stderr.String())
		}
	}
}

func TestMissingConfig(t *testing.T) {
	if code := run([]string{"plugin"}, &bytes.Buffer{}, &bytes.Buffer{}); code != ErrMissingConfig {
		t.Errorf("exit %d, want %d", code, ErrMissingConfig)
	}
}
