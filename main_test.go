package exregexp

import (
	"fmt"
	"regexp"
	"testing"
)

func TestReplaceAllStringSubmatchFunc(t *testing.T) {
	rx, err := regexp.Compile(`\[\[(?:([^\|\]]+)\|)?(.+?)\]\]`)
	if err != nil {
		t.Fatalf("test regexp error: %s", err.Error())
	}

	result := ReplaceAllStringSubmatchFunc(rx,
		"foo[[LINKED TEXT|URL]]bar",
		func(s []string) string {
			return fmt.Sprintf(`<a href="%s">%s</a><!-- %s -->`, s[2], s[1], s[0])
		})

	expect := `foo<a href="URL">LINKED TEXT</a><!-- [[LINKED TEXT|URL]] -->bar`
	if result != expect {
		t.Fatalf("(1) expect `%s`, but `%s`", expect, result)
	}

	result = ReplaceAllStringSubmatchFunc(rx,
		"foo[[LINKED TEXT|URL]]bar[[linked text|url]]qux",
		func(s []string) string {
			return fmt.Sprintf(`<a href="%s">%s</a><!-- %s -->`, s[2], s[1], s[0])
		})
	expect = `foo<a href="URL">LINKED TEXT</a><!-- [[LINKED TEXT|URL]] -->bar<a href="url">linked text</a><!-- [[linked text|url]] -->qux`
	if result != expect {
		t.Fatalf("(2) expect `%s`, but `%s`", expect, result)
	}

	result = ReplaceAllStringSubmatchFunc(rx,
		"A[[1|2]]B[[3|4]]C[[5|6]]D",
		func(s []string) string {
			return fmt.Sprintf(`[%s](%s)<!-- %s -->`, s[1], s[2], s[0])
		})
	expect = `A[1](2)<!-- [[1|2]] -->B[3](4)<!-- [[3|4]] -->C[5](6)<!-- [[5|6]] -->D`
	if result != expect {
		t.Fatalf("(2) expect `%s`, but `%s`", expect, result)
	}
}
