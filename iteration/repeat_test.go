package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assert := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	}

	t.Run("Repeat positive number of times", func(t *testing.T) {
		assert(t, Repeat("a", 5), "aaaaa")
	})

	t.Run("Repeat 0 times", func(t *testing.T) {
		assert(t, Repeat("a", 0), "")
	})

	t.Run("Repeat neg times", func(t *testing.T) {
		assert(t, Repeat("a", -20), "a")
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
