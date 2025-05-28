package exregexp

import (
	"bytes"
	"strings"
)

type RegexpBin interface {
	FindAllSubmatchIndex([]byte, int) [][]int
}

func ReplaceAllSubmatchFunc(rx RegexpBin, input []byte, f func([][]byte) []byte) []byte {
	var result bytes.Buffer

	matches := rx.FindAllSubmatchIndex(input, -1)
	lastIndex := 0
	for _, match := range matches {
		result.Write(input[lastIndex:match[0]])

		sub := make([][]byte, len(match))
		for i := 0; i*2 < len(match); i++ {
			if from := match[i*2]; from < 0 {
				sub[i] = []byte{}
			} else if to := match[i*2+1]; to < 0 {
				sub[i] = []byte{}
			} else {
				sub[i] = input[from:to]
			}
		}
		result.Write(f(sub))

		lastIndex = match[1]
	}
	result.Write(input[lastIndex:])
	return result.Bytes()
}

type RegexpStr interface {
	FindAllStringSubmatchIndex(string, int) [][]int
}

// ReplaceAllStringSubmatchFunc applies a replacement function to all
// matches of the regular expression in the input string. The replacement
// function receives a slice of submatches, where the first element is the
// entire match and subsequent elements are the capturing groups.
func ReplaceAllStringSubmatchFunc(rx RegexpStr, input string, f func([]string) string) string {
	var result strings.Builder

	matches := rx.FindAllStringSubmatchIndex(input, -1)
	lastIndex := 0
	for _, match := range matches {
		result.WriteString(input[lastIndex:match[0]])

		sub := make([]string, len(match))
		for i := 0; i*2 < len(match); i++ {
			if from := match[i*2]; from < 0 {
				sub[i] = ""
			} else if to := match[i*2+1]; to < 0 {
				sub[i] = ""
			} else {
				sub[i] = input[from:to]
			}
		}
		result.WriteString(f(sub))

		lastIndex = match[1]
	}
	result.WriteString(input[lastIndex:])
	return result.String()
}
