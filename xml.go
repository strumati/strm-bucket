package main

import (
	"fmt"
	"strings"
)

func Xml(name string, params map[string]string, elmts []string) string {
	var attrs []string
	for k, v := range params {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, k, v))
	}

	if len(attrs) > 0 {
		return fmt.Sprintf(
			"<%s %s>%s</%s>",
			name,
			strings.Join(attrs, " "),
			strings.Join(elmts, ""),
			name,
		)
	}

	return fmt.Sprintf(
		"<%s>%s</%s>",
		name,
		strings.Join(elmts, ""),
		name,
	)
}
