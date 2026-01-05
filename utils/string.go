package utils

import (
	"strings"
	"unicode"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/slices"
)

func IsDigit(content string) bool {
	characters := strings.Split(content, "")

	return slices.Every(characters, func(s string, i int) bool {
		return unicode.IsDigit([]rune(s)[0])
	})
}
