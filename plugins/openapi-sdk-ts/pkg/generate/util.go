package generate

import (
	"strconv"
	"strings"
	"unicode"
)

func toPascalCase(s string) string {
	parts := splitWords(s)
	var b strings.Builder
	for _, p := range parts {
		if len(p) == 0 {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]) + strings.ToLower(p[1:]))
	}
	return b.String()
}

func toCamelCase(s string) string {
	p := toPascalCase(s)
	if len(p) == 0 {
		return p
	}
	return strings.ToLower(p[:1]) + p[1:]
}

func splitWords(s string) []string {
	var words []string
	var current strings.Builder
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			if current.Len() > 0 {
				words = append(words, current.String())
				current.Reset()
			}
			continue
		}
		if unicode.IsUpper(r) && i > 0 {
			prev := rune(s[i-1])
			if unicode.IsLower(prev) || prev == '_' || prev == '-' {
				if current.Len() > 0 {
					words = append(words, current.String())
					current.Reset()
				}
			}
		}
		current.WriteRune(r)
	}
	if current.Len() > 0 {
		words = append(words, current.String())
	}
	return words
}

// isIdent reports whether s is a valid ASCII TypeScript identifier.
func isIdent(s string) bool {
	if s == "" {
		return false
	}
	for i, r := range s {
		ok := r == '_' || r == '$' || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (i > 0 && r >= '0' && r <= '9')
		if !ok {
			return false
		}
	}
	return true
}

// tsReserved lists words that cannot be used as a parameter or function name,
// plus local names the generated client methods use themselves.
var tsReserved = map[string]bool{
	"break": true, "case": true, "catch": true, "class": true, "const": true, "continue": true,
	"debugger": true, "default": true, "delete": true, "do": true, "else": true, "enum": true,
	"export": true, "extends": true, "false": true, "finally": true, "for": true, "function": true,
	"if": true, "import": true, "in": true, "instanceof": true, "new": true, "null": true,
	"return": true, "super": true, "switch": true, "this": true, "throw": true, "true": true,
	"try": true, "typeof": true, "var": true, "void": true, "while": true, "with": true,
	"let": true, "static": true, "yield": true, "await": true, "implements": true, "interface": true,
	"package": true, "private": true, "protected": true, "public": true, "arguments": true, "eval": true,
	// locals of generated methods
	"body": true, "url": true, "query": true, "resp": true,
}

// tsSafeName makes a camelCase name usable as an identifier: it never starts
// with a digit, is never empty and never collides with a reserved word.
func tsSafeName(s string) string {
	if s == "" {
		return "_"
	}
	if s[0] >= '0' && s[0] <= '9' {
		s = "_" + s
	}
	if tsReserved[s] {
		s += "_"
	}
	return s
}

// tsPropName renders an object property name, quoting it when it is not a plain identifier.
func tsPropName(name string) string {
	if isIdent(name) {
		return name
	}
	return strconv.Quote(name)
}

// tsTypeName turns a schema name into a TypeScript type name. Names that are
// already identifiers keep their spelling (first letter upper-cased); others
// (`foo-bar`, `meta/root`) are split on the invalid characters and re-joined.
func tsTypeName(name string) string {
	if name == "" {
		return ""
	}
	if !isIdent(name) {
		var b strings.Builder
		for _, w := range splitNonIdent(name) {
			b.WriteString(strings.ToUpper(w[:1]) + w[1:])
		}
		name = b.String()
		if name == "" {
			return "Unnamed"
		}
		if name[0] >= '0' && name[0] <= '9' {
			name = "T" + name
		}
	}
	r := []rune(name)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func splitNonIdent(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return !(r == '_' || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'))
	})
}

func refToTypeName(ref string) string {
	parts := strings.Split(ref, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}
