package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Ivan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Ivan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellos(t *testing.T) {
	names := []string{"Ivan", "Vlad", "Sergey"}
	messages, err := Hellos(names)
	if err != nil {
		t.Fatalf(`Error found: %v`, err)
	}
	for name, msg := range messages {
		want := regexp.MustCompile(`\b` + name + `\b`)
		if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Hello(%v) = %q, %v, want match for %#q, nil`, name, msg, err, want)
		}
	}
}
