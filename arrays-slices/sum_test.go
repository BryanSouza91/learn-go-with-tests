package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("set of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want, numbers)
	})

	t.Run("set of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want, numbers)
	})
}

// Fails with "got [3 10 9] want [3 10 9]"
// func TestSumAll(t *testing.T) {
// 	got := SumAll([]int{1, 2}, []int{1, 2, 3, 4}, []int{1, 3, 5})
// 	want := []int{3, 10, 9}

// 	if reflect.DeepEqual(got, want) {
// 		t.Errorf("got %+v want %+v", got, want)
// 	}
// }

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sums the tails of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{1, 2, 3, 4}, []int{1, 3, 5})
		want := []int{2, 9, 8}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)

	})

}

