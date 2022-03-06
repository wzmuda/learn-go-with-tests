package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	assert := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("simple case", func(t *testing.T) {
		assert(t, SumAllTails([]int{1, 2}, []int{0, 9}), []int{2, 9})
	})

	t.Run("empty slice", func(t *testing.T) {
		assert(t, SumAllTails([]int{}, []int{3, 4, 5}), []int{0, 9})
	})
}
