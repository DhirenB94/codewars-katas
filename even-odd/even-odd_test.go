package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	t.Run("Even number", func(t *testing.T) {

	got := EvenOrOdd(20)
	want:= "Even"
	assertCorrectMessage(t, got, want)

	})
	t.Run("Odd number", func(t *testing.T) {

		got := EvenOrOdd(9)
		want:= "Odd"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	 t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}