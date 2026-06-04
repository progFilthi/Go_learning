package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := hello("FBI")
		want := "Hello, FBI"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("When the name is empty", func(t *testing.T) {
		got := hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
