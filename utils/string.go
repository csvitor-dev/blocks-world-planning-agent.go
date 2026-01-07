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

func Center(s string, width int) string {
	if len(s) >= width {
		return s
	}
	padding := width - len(s)
	left := padding / 2
	right := padding - left

	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}
