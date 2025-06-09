package test

import (
	mainapp "learn-go-with-tests/src"
	"testing"
)

func TestLoop(t *testing.T) {
	t.Run("Call loop with a and get five a", func(t *testing.T) {
		got := mainapp.Loop("a")
		want := "aaaaa"

		if got != want {
			t.Errorf("Wanted '%s' got '%s'", want, got)
		}
	})
}
