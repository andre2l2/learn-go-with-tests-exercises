package test

import (
	mainapp "learn-go-with-tests/src"
	"testing"
)

func TestAddNumber(t *testing.T) {
	t.Run("Add 2 + 2 = 4", func(t *testing.T) {
		got := mainapp.AddNumber(2, 2)
		want := 4

		if got != want {
			t.Errorf("Expect '%d', recive '%d'", want, got)	
		}
	})
}
