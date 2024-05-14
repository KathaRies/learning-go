package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Merry"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name, 10)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Merry") = %q,%v want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("", 5)
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
