package main

import "testing"

func TestHello(t *testing.T) { //I learned that test functions need to start with a capital T!!
	got := hello("tom")
	want := "Hello, tom!"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
