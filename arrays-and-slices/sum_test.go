package main

import "testing"

func TestSum(t *testing.T) {
	assert := func(t testing.TB, got, want int, array []int) {
		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, array)
		}
	}

	numbers := []int{1, 2, 3, 4, 5}
	t.Run("first 5 naturals", func(t *testing.T) {
		assert(t, Sum(numbers), 15, numbers)
	})

	numbers = []int{1, -1}
	t.Run("negative pair", func(t *testing.T) {
		assert(t, Sum(numbers), 0, numbers)
	})

	numbers = []int{21, 37, 666}
	t.Run("3 random ints", func(t *testing.T) {
		assert(t, Sum(numbers), 724, numbers)
	})
}
