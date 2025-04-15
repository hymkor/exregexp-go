//go:build example

package main

import (
	"fmt"
	"regexp"

	"github.com/hymkor/exregexp-go"
)

func main() {
	re := regexp.MustCompile(`\b([a-zA-Z]+)(\d+)\b`)
	input := "example123 test456 hello789"
	output := exregexp.ReplaceAllStringSubmatchFunc(re, input, func(submatches []string) string {
		return fmt.Sprintf("%s(%s)", submatches[1], submatches[2])
	})

	fmt.Println(output)
}
