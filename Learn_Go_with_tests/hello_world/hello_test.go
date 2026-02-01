package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Umama")
	want := "Hello, Umama"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
