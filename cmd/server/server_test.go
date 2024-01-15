package main

import (
	"testing"
)

func TestRandomQuote(t *testing.T) {
	quote := randomQuote()
	if len(quote) == 0 {
		t.Errorf("randomQuote returned an empty string")
	}
	found := false
	for _, q := range quotes {
		if q == quote {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("randomQuote returned a string not in the quotes array")
	}
}
