package openapi

import (
	"strings"
)


func concatenateDescription(l string, r string) string {
	var sb strings.Builder
	if l != "" {
		sb.WriteString(l)
	}
	if r != "" {
		if sb.Len() > 0 {
			sb.WriteString(": ")
		}
		sb.WriteString(r)
	}
	return sb.String()
}


func escapeTabs(s string) string {
	return strings.ReplaceAll(s, "\t", "\\t")
}

func escapeNewLines(s string) string {
	return strings.ReplaceAll(s, "\n", "\\n")
}
