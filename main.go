package exregexp

import (
	"regexp"
)

func ReplaceAllStringSubmatchFunc(rx *regexp.Regexp, s string, f func([]string) string) string {
	matches := rx.FindAllStringSubmatchIndex(s, -1)
	for len(matches) > 0 {
		last := matches[len(matches)-1]
		matches = matches[:len(matches)-1]

		ss := make([]string, 0, len(last)/2)
		for i := 0; i < len(last); i += 2 {
			ss = append(ss, s[last[i]:last[i+1]])
		}
		s = s[:last[0]] + f(ss) + s[last[1]:]
	}
	return s
}
