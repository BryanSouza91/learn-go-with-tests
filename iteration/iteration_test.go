package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("x", 3)
	expected := "xxx"

	assertCorrectMessage(t, repeated, expected)
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %s but got %s", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// This is an example of how to use the Repeat func
func ExampleRepeat() {
	repeated := Repeat("j", 3)
	fmt.Println(repeated)
	// Output: jjj
}
