package arrays

import "testing"

func TestSum(t *testing.T) {

	//A slice is not fixed sized while an array is fixed sized.

	t.Run("Testing a : slice", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)

		want := 15

		if got != want {
			t.Errorf("got %d, want %d, numbers %v", got, want, nums)
		}
	})

	t.Run("Testing an : array", func(t *testing.T) {

		nums := [3]int{1, 2, 3}

		got := Sum(nums[:])

		want := 6

		if got != want {
			t.Errorf("got %d, want %d, numbers %v", got, want, nums)
		}

	})

}
