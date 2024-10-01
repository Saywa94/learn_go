package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collections of 5 numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(numbers)

		if want != got {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})

	t.Run("Collections of any size", func(t *testing.T) {

		numbers := []int{3, 2, 5}
		want := 10
		got := Sum(numbers)

		if want != got {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("Sum totals of multiple collections", func(t *testing.T) {

		want := []int{10, 6, 20}
		got := SumAll([]int{3, 2, 5}, []int{2, 2, 2}, []int{1, 9, 10})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

}
func TestSumAllTails(t *testing.T) {
	t.Run("Sum tails of multiple collections", func(t *testing.T) {

		want := []int{7, 4, 19}
		got := SumAllTails([]int{3, 2, 5}, []int{2, 2, 2}, []int{1, 9, 10})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}

	})
	t.Run("Sum tails of multiple collections, passing empty slice", func(t *testing.T) {

		want := []int{}
		got := SumAllTails()

		if len(got) > 0 {
			t.Errorf("got %v, want %v", got, want)
		}

	})

}
