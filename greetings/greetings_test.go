package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Sonali"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Sonali")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Sonali") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
