// Test rules:
// file must be named xxx_test.go
package main

import "testing"

// function must be named TestXXX
// it takes always one t *testing.T argument
func TestHello(t *testing.T) {
	// for some reason this can be classic function but outside
	// of TestHello()
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			// %q wraps values in "double quotes"
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello, World on empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})
}
