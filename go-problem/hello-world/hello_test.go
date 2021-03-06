package main

import (
	"bytes"
	"testing"
)

// Define a function HelloWorld(string) string.
//
// Also define a testVersion with a value that matches
// the targetTestVersion here.

const targetTestVersion = 2

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestHelloWorld(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, input, got, want string) {
		t.Helper()
		if got != want {
			t.Fatalf("HelloWorld(%s) = %v, want %v", input, got, want)
		}
	}

	t.Run("in French", func(t *testing.T) {
		input := "Elodie"
		lang := french
		got := Helloworld(input, lang)
		want := "Bonjour, Elodie!"
		assertCorrectMessage(t, input, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		input := "Elodie"
		lang := spanish
		got := Helloworld(input, lang)
		want := "Hola, Elodie!"
		assertCorrectMessage(t, input, got, want)
	})

	t.Run("empty name", func(t *testing.T) {
		input := ""
		lang := ""
		got := Helloworld(input, lang)
		assertCorrectMessage(t, input, got, "Hello, World!")
	})

	t.Run("in English", func(t *testing.T) {
		input := "Gopher"
		lang := ""
		got := Helloworld(input, lang)
		assertCorrectMessage(t, input, got, "Hello, Gopher!")
	})

}
