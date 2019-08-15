package main

import (
	"testing"
	"fmt"
)

func TestGreeting(t *testing.T) {

	got := Greeting("George")

	want := "Hello George"

	if got != want {

		t.Fatalf("Expected %q, got %q", want, got)

	} else {
		fmt.Println("Expected ", want, "got", got)
	}

}
