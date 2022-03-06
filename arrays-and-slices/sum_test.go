package main

import "testing"

func TestSum(t *testing.T) {
	assert := func(t testing.TB, got, want int, array [5]int) {
		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, array)
		}
	}

	numbers := [5]int{1, 2, 3, 4, 5}
	t.Run("first 5 naturals", func(t *testing.T) {
		assert(t, Sum(numbers), 15, numbers)
	})

	numbers = [5]int{1, -1}
	t.Run("negative", func(t *testing.T) {
		assert(t, Sum(numbers), 0, numbers)
	})
}
