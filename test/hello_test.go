package test

import (
	mainapp "learn-go-with-tests/src"
	"testing"
)

func TestCustomHello(t *testing.T) {
	t.Run("Say hello for peoples", func(t *testing.T) {
		got := mainapp.CustomHello("Andre", "")
		want := "Hello, Andre!"

		if got != want {
			t.Errorf("The result '%s' expect '%s'", got, want)
		}
	})

	t.Run("Say hello for peoples with default", func(t *testing.T) {
		got := mainapp.CustomHello("", "")
		want := "Hello!"

		if got != want {
			t.Errorf("The result '%s' expect '%s'", got, want)
		}
	})

	t.Run("Say hello in spanish", func(t *testing.T) {
		got := mainapp.CustomHello("Andre", "spanish")
		want := "Hola, Andre!"

		if got != want {
			t.Errorf("The result '%s' expect '%s'", got, want)
		}
	})

	t.Run("Say hello in french", func(t *testing.T) {
		got := mainapp.CustomHello("Andre", "french")
		want := "Bonjour, Andre!"

		if got != want {
			t.Errorf("The result '%s', expect '%s'", got, want)
		}
	})
}
