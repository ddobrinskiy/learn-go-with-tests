package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {

		got := Hello("David", "")
		want := "Hello, David"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' for an empty string", func(t *testing.T) {

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
