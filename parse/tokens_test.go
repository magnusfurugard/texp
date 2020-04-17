package parse

import (
	"reflect"
	"testing"
)

var cases = []struct {
	query    string
	tokens   *[]string
	expected []string
}{
	{"@a+@a=@b", &[]string{"@a=1", "@b=2"}, []string{"1+1=2"}},
	{"hello @person", &[]string{"@person=a", "@person=b"}, []string{"hello a", "hello b"}},
	{"hello @person", &[]string{"@person=a", "@person=b", "@person=c", "@person=d"}, []string{"hello a", "hello b", "hello c", "hello d"}},
	{"@a@a@b", &[]string{"@a=1", "@a=2", "@b=3"}, []string{"113", "223"}},
	{"@a@a@b@b", &[]string{"@a=1", "@a=2", "@b=3", "@b=4"}, []string{"1133", "1144", "2233", "2244"}},
}

func TestTokens(t *testing.T) {
	for _, c := range cases {
		_, r := Tokens(c.query, c.tokens)
		if !reflect.DeepEqual(r, c.expected) {
			t.Errorf("expected: %v, got: %v", c.expected, r)
		}
	}
}
