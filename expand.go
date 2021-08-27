package goutils

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Expand replaces ${var} in the template string base on the mapping function.
//
// It's like os.Expand but with following differences:
//   1. strict syntax checks, return the string replaces as many as we can, and an syntax error.
//   2. compatible with bash like ${var:=default} syntax.
//   3. $var outputs as is.
func Expand(s string, mapping func(varName string, default_ string) string) (string, error) {
	if len(s) == 0 {
		return s, nil
	}
	var sb strings.Builder
	if !utf8.FullRuneInString(s) {
		return s, fmt.Errorf("invalid utf8 string: %s", s)
	}
	var err error
	i := 0
	for j := 0; j < len(s); {
		r, cursor := nextUTF8Character(s, j)
		size := cursor - j
		if r == '$' && j+1 < len(s) {
			sb.WriteString(s[i:j])
			name, w := getShellName(s[j+1:])
			if name == "" && w > 0 {
				// Encountered invalid syntax; eat the characters.
				if err == nil {
					err = fmt.Errorf("bad syntax at %d: %s", j, s)
				}
				sb.WriteString(s[j : j+w+1])
			} else if name == "" {
				// Valid syntax, but $ was not followed by a
				// name. Leave the dollar character untouched.
				sb.WriteByte(s[j])
			} else {
				sb.WriteString(mapping(parseVar(name)))
			}
			j += w
			i = j + 1
		}
		j += size
	}
	if sb.Len() == 0 {
		return s, err
	}
	sb.WriteString(s[i:])
	return sb.String(), err
}

func nextUTF8Character(s string, cursor int) (rune, int) {
	r, size := utf8.DecodeRuneInString(s[cursor:])
	return r, cursor + size
}

func getShellName(s string) (string, int) {
	r, cursor := nextUTF8Character(s, 0)
	if r != '{' {
		return "", 0
	}
	// Scan to closing brace
	for i := cursor; i < len(s); {
		r, cur := nextUTF8Character(s, i)
		if r == '}' {
			if i == 1 {
				return "", 2 // Bad syntax; eat "${}"
			}
			return s[1:i], cur
		}
		if r == '{' || r == '$' {
			break
		}
		i = cur
	}
	return "", 1 // Bad syntax; eat "${"
}

func parseVar(va string) (string, string) {
	a := strings.SplitN(va, ":=", 2)
	if len(a) < 2 {
		return va, ""
	}
	return a[0], a[1]
}
