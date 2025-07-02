package string_utils

import "strings"

func ToCamelCase(input string, sep rune, upperFirst bool) string {
	parts := strings.Split(input, string(sep))
	for i, part := range parts {
		if i == 0 && !upperFirst {
			parts[i] = strings.ToLower(part)
		} else {
			parts[i] = strings.Title(strings.ToLower(part))
		}
	}
	return strings.Join(parts, "")
}
