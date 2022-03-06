package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	assert := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("xxx", func(t *testing.T) {
		assert(t, SumAll([]int{1, 2}, []int{0, 9}), []int{3, 9})
	})
}
