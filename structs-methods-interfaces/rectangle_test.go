package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	assert := func(t testing.TB, got, want float64) {
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	t.Run("", func(t *testing.T) {
		assert(t, Perimeter(Rectangle{10.0, 10.}), 40.0)
	})
}
