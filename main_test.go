package exregexp

import (
	"fmt"
	"bytes"
	"regexp"
	"testing"
)

func concat(s ...[]byte) []byte {
	r := []byte{}
	for _, s1 := range s {
		r = append(r, s1...)
	}
	return r
}

func TestReplaceAllSubmatchFunc(t *testing.T) {
	rx, err := regexp.Compile(`\[\[(?:([^\|\]]+)\|)?(.+?)\]\]`)
	if err != nil {
		t.Fatalf("test regexp error: %s", err.Error())
	}

	result := ReplaceAllSubmatchFunc(rx,
		[]byte("foo[[LINKED TEXT|URL]]bar"),
		func(s [][]byte) []byte {
			return []byte(fmt.Sprintf(`<a href="%s">%s</a><!-- %s -->`, string(s[2]), string(s[1]), string(s[0])))
		})

	expect := []byte(`foo<a href="URL">LINKED TEXT</a><!-- [[LINKED TEXT|URL]] -->bar`)
	if !bytes.Equal(result, expect) {
		t.Fatalf("(1) expect `%v`, but `%v`", expect, result)
	}

	result = ReplaceAllSubmatchFunc(rx,
		[]byte("foo[[LINKED TEXT|URL]]bar[[linked text|url]]qux"),
		func(s [][]byte) []byte {
			return []byte(fmt.Sprintf(`<a href="%s">%s</a><!-- %s -->`, string(s[2]), string(s[1]), string(s[0])))
		})
	expect = []byte(`foo<a href="URL">LINKED TEXT</a><!-- [[LINKED TEXT|URL]] -->bar<a href="url">linked text</a><!-- [[linked text|url]] -->qux`)
	if !bytes.Equal(result, expect) {
		t.Fatalf("(2) expect `%v`, but `%v`", expect, result)
	}

	result = ReplaceAllSubmatchFunc(rx,
		[]byte("A[[1|2]]B[[3|4]]C[[5|6]]D"),
		func(s [][]byte) []byte {
			return []byte(fmt.Sprintf(`[%s](%s)<!-- %s -->`, string(s[1]), string(s[2]), string(s[0])))
		})
	expect = []byte(`A[1](2)<!-- [[1|2]] -->B[3](4)<!-- [[3|4]] -->C[5](6)<!-- [[5|6]] -->D`)
	if !bytes.Equal(result, expect) {
		t.Fatalf("(2) expect `%v`, but `%v`", expect, result)
	}

	rx, err = regexp.Compile(`([a-zA-Z]+)([0-9])?`)
	if err != nil {
		t.Fatalf("test regexp error: %s", err.Error())
	}
	result = ReplaceAllSubmatchFunc(rx,
		[]byte("aiueo1"),
		func(s [][]byte) []byte {
			return concat(s[1], []byte{'-'}, s[2])
		})
	expect = []byte("aiueo-1")
	if !bytes.Equal(result, expect) {
		t.Fatalf("(3) expect `%v`, but `%v`", expect, result)
	}
	result = ReplaceAllSubmatchFunc(rx,
		[]byte("aiueo"),
		func(s [][]byte) []byte {
			return concat(s[1], []byte{'-'}, s[2])
		})
	expect = []byte("aiueo-")
	if !bytes.Equal(result, expect) {
		t.Fatalf("(4) expect `%v`, but `%v`", expect, result)
	}
}

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

	rx, err = regexp.Compile(`([a-zA-Z]+)([0-9])?`)
	if err != nil {
		t.Fatalf("test regexp error: %s", err.Error())
	}
	result = ReplaceAllStringSubmatchFunc(rx,
		"aiueo1",
		func(s []string) string {
			return s[1] + "-" + s[2]
		})
	expect = "aiueo-1"
	if result != expect {
		t.Fatalf("(3) expect `%s`, but `%s`", expect, result)
	}
	result = ReplaceAllStringSubmatchFunc(rx,
		"aiueo",
		func(s []string) string {
			return s[1] + "-" + s[2]
		})
	expect = "aiueo-"
	if result != expect {
		t.Fatalf("(4) expect `%s`, but `%s`", expect, result)
	}
}
