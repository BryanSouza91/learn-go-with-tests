package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("x")
	expected := "xxx"

	assertCorrectMessage(t, repeated, expected)
}
func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %s but got %s", expected, repeated)
	}

}
