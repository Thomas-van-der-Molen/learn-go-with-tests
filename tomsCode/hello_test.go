package main

import "testing"

func TestHello(t *testing.T) { //I learned that test functions need to start with a capital T!!
	t.Run("testing with a person's name", func(t *testing.T) {
		got := hello("tom")
		want := "Hello, tom!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("testing with no input", func(t *testing.T) {
		got := hello("")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
