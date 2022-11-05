package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Bryan")
	want := "Hello, Bryan!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
