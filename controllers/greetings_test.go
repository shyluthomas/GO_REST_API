package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "shylu"
	prop := Props{Name: name}
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(prop)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("shylu") = %q, %v, want match for %q`, msg, err, want)
	}
}
