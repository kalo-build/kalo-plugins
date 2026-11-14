package generate

import (
	"fmt"
	"io"
	"regexp"
)

var openapi31 = regexp.MustCompile(`(?m)^\s*"?openapi"?\s*:\s*["']?3\.1`)

// warnIfOpenAPI31 writes one line to w when the document declares OpenAPI 3.1. The
// generator reads the 3.0 subset: 3.1-only keywords (null in a type list,
// exclusiveMinimum as a number, ...) are ignored without any other message.
func warnIfOpenAPI31(w io.Writer, spec []byte) {
	if w != nil && openapi31.Match(spec) {
		fmt.Fprintln(w, "warning: openapi 3.1 input: only OpenAPI 3.0 keywords are read, 3.1-only keywords are ignored")
	}
}
