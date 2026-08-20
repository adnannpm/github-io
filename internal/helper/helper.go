package helper

import "strings"

func PrintableValue(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "-"
	}

	return strings.NewReplacer("\r", " ", "\n", " ", "\t", " ").Replace(value)
}