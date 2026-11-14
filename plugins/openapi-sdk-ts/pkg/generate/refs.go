package generate

import (
	"fmt"
	"sort"
	"strings"
)

const schemaRefPrefix = "#/components/schemas/"

// CheckRefs returns an error when the document holds a schema $ref the generators
// would turn into a type name without defining it: anything that is not
// #/components/schemas/<Name>, or a <Name> that components.schemas does not define.
// It reports the first such $ref and how many there are, in one line.
func CheckRefs(doc *Document) error {
	defined := make(map[string]bool, len(doc.Schemas))
	for _, ns := range doc.Schemas {
		defined[ns.Name] = true
	}

	var first string
	count := 0
	check := func(ref string) {
		if ref == "" {
			return
		}
		name := strings.TrimPrefix(ref, schemaRefPrefix)
		if strings.HasPrefix(ref, schemaRefPrefix) && !strings.Contains(name, "/") && defined[name] {
			return
		}
		if count == 0 {
			first = ref
		}
		count++
	}

	var walkSchema func(s Schema)
	walkRef := func(sr *SchemaRef) {
		if sr == nil {
			return
		}
		check(sr.Ref)
		if sr.Schema != nil && sr.Ref == "" {
			walkSchema(*sr.Schema)
		}
	}
	walkSchema = func(s Schema) {
		for _, p := range s.Properties {
			check(p.Ref)
			if p.Ref == "" {
				walkSchema(p.Schema)
			}
		}
		walkRef(s.Items)
	}

	for _, pe := range doc.Paths {
		for _, op := range pe.Operations {
			for _, p := range op.Parameters {
				if p.Schema != nil {
					walkSchema(*p.Schema)
				}
			}
			walkRef(op.RequestBody)
			codes := make([]string, 0, len(op.Responses))
			for c := range op.Responses {
				codes = append(codes, c)
			}
			sort.Strings(codes)
			for _, c := range codes {
				walkRef(op.Responses[c])
			}
		}
	}
	for _, ns := range doc.Schemas {
		walkSchema(ns.Schema)
	}

	if count == 0 {
		return nil
	}
	return fmt.Errorf("unsupported $ref %q (%d such $refs in the schemas read): only $refs to #/components/schemas/<Name> are supported", first, count)
}
