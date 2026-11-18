package generate_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-go/pkg/generate"
	"github.com/stretchr/testify/require"
)

// denseSpec has n schemas, each with an optional $ref property to every schema, and
// one GET returning S0. Expanding every path through that graph is factorial in n
// (n=8 ran the 2 GB runner out of memory); a $ref'd property is never expanded.
func denseSpec(n int) []byte {
	var schemas []string
	for i := 0; i < n; i++ {
		var props []string
		for j := 0; j < n; j++ {
			props = append(props, fmt.Sprintf(`"p%d":{"$ref":"#/components/schemas/S%d"}`, j, j))
		}
		schemas = append(schemas, fmt.Sprintf(`"S%d":{"type":"object","properties":{%s}}`, i, strings.Join(props, ",")))
	}
	return []byte(fmt.Sprintf(`{"openapi":"3.0.3","info":{"title":"dense","version":"1"},
"paths":{"/s":{"get":{"operationId":"getS","responses":{"200":{"description":"ok",
"content":{"application/json":{"schema":{"$ref":"#/components/schemas/S0"}}}}}}}},
"components":{"schemas":{%s}}}`, strings.Join(schemas, ",")))
}

func TestLinear_DenselyLinkedSchemas(t *testing.T) {
	start := time.Now()
	doc, err := generate.ParseSpec(denseSpec(40))
	require.NoError(t, err)
	_, err = generate.GenerateTypes(doc, "sdk")
	require.NoError(t, err)
	_, err = generate.GenerateClient(doc, "sdk")
	require.NoError(t, err)
	require.Less(t, time.Since(start), 10*time.Second)
}

// A {data: $ref ItemList} wrapper, where ItemList is an array of $ref Item, returns
// the item list: the shallow schema of a $ref'd property keeps its items' ref.
func TestLinear_WrappedRefArrayReturn(t *testing.T) {
	spec := `{"openapi":"3.0.3","info":{"title":"w","version":"1"},
"paths":{"/items":{"get":{"operationId":"listItems","responses":{"200":{"description":"ok",
"content":{"application/json":{"schema":{"type":"object","properties":{"data":{"$ref":"#/components/schemas/ItemList"}}}}}}}}}},
"components":{"schemas":{"Item":{"type":"object","properties":{"id":{"type":"string"}}},
"ItemList":{"type":"array","items":{"$ref":"#/components/schemas/Item"}}}}}`
	doc, err := generate.ParseSpec([]byte(spec))
	require.NoError(t, err)
	client, err := generate.GenerateClient(doc, "sdk")
	require.NoError(t, err)
	require.Regexp(t, `\[\]Item\b`, client)
	require.NotContains(t, client, "ItemList")
}
