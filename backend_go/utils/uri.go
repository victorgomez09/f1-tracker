package utils

import (
	"net/url"
	"strings"
)

func UrlEncoded(str string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(
							url.QueryEscape(strings.ReplaceAll(str, " ", "%20")),
							"%2520",
							"%20",
						),
						"%27",
						"'",
					),
					"%28",
					"(",
				),
				"%29",
				")",
			),
			"%21",
			"!",
		),
		"%2A",
		"*",
	)
}