package exregexp

import (
	"regexp"
	"strings"
)

// ReplaceAllStringSubmatchFunc applies a replacement function to all
// matches of the regular expression in the input string. The replacement
// function receives a slice of submatches, where the first element is the
// entire match and subsequent elements are the capturing groups.
func ReplaceAllStringSubmatchFunc(rx *regexp.Regexp, input string, f func([]string) string) string {
	var result strings.Builder

	matches := rx.FindAllStringSubmatchIndex(input, -1)
	lastIndex := 0
	for _, match := range matches {
		result.WriteString(input[lastIndex:match[0]])

		sub := make([]string, len(match))
		for i := 0; i*2 < len(match); i++ {
			sub[i] = input[match[i*2]:match[i*2+1]]
		}
		result.WriteString(f(sub))

		lastIndex = match[1]
	}
	result.WriteString(input[lastIndex:])
	return result.String()
}
