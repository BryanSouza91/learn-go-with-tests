package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Bryan")

		got := buffer.String()
		want := "Hello, Bryan!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
