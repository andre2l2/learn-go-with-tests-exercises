package test

import (
	mainapp "learn-go-with-tests/src"
	"testing"
)

func TestHello(t *testing.T) {
	result := mainapp.Hello()
	expect := "Hello!"

	if result != expect {
		t.Errorf("The result '%s' expect '%s'", result, expect)
	}
}

func TestCustomHello(t *testing.T) {
	result := mainapp.CustomHello("Andre")
	expect := "Hello, Andre!"

	if result != expect {
		t.Errorf("The result '%s' expect '%s'", result, expect)
	}
}
