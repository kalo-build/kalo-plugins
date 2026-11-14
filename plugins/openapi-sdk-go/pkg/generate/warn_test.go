package generate

import (
	"strings"
	"testing"
)

func TestWarnIfOpenAPI31(t *testing.T) {
	cases := map[string]bool{
		"openapi: 3.1.0\ninfo: {}":          true,
		"openapi: \"3.1.1\"\n":              true,
		"{\n  \"openapi\": \"3.1.0\"\n}":    true,
		"openapi: 3.0.3\n":                  false,
		"info:\n  description: openapi: 3.1": false,
	}
	for in, want := range cases {
		var b strings.Builder
		warnIfOpenAPI31(&b, []byte(in))
		if got := b.Len() > 0; got != want {
			t.Errorf("%q: warned=%v, want %v", in, got, want)
		}
	}
}
