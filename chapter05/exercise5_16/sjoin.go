package variadic

import "strings"

func StringJoin(separator string, list ...string) string {
	return strings.Join(list, separator)
}
