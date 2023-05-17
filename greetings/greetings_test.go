package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Go"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Go")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Go") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
