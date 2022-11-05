package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bryan")
		want := "Hello, Bryan!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, wonderful world!' when an empty string is entered", func(t *testing.T) {
		got := Hello("")
		want := "Hello, wonderful world!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
