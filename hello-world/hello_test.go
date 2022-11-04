package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, wonderful world!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
