package arrays_and_slices

import "testing"

func TestSum(t *testing.T)  {

	t.Run("sum array[1,2,3,4,5] should be 15", func(t *testing.T) {
		numbers := [5]int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("sum array[5,6,7,8,9] should be 35", func(t *testing.T) {
		numbers := [5]int{5,6,7,8,9}

		got := Sum(numbers)
		want := 35

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
