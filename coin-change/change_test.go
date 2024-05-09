package coinchange

import (
	"fmt"
	"testing"
)

func isEq(t testing.TB, a, b []int) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

var cases = []struct {
	Due   int
	Coins []int
}{
	{Due: 42, Coins: []int{25, 10, 5, 1, 1}},
	{Due: 12, Coins: []int{10, 1, 1}},
	{Due: 29, Coins: []int{25, 1, 1, 1, 1}},
	{Due: 99, Coins: []int{25, 25, 25, 10, 10, 1, 1, 1, 1}},
}

func TestChange(t *testing.T) {
	for _, test := range cases {
		t.Run(
			fmt.Sprintf("%v gets converted to %v", test.Due, test.Coins),
			func(t *testing.T) {
				got := Change(test.Due)
				want := test.Coins
				eq := isEq(t, got, want)
				if !eq {
					t.Errorf("got %v want %v", got, want)
				}
			})
	}

}
