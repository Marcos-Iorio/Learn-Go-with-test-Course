package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Marcos", "English")
		want := "Hello, Marcos"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'hello, world' when no argument is passed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Says hello in spanish", func(t *testing.T) {
		got := Hello("Marcos", "Spanish")
		want := "Hola, Marcos"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Says hello in French", func(t *testing.T) {
		got := Hello("Marcos", "French")
		want := "Bonjour, Marcos"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
