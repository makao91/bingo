package main

import (
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestGreeting(t *testing.T) {
	out, in := makeUI()
	test.Type(in, "Andy")
	if out.Text != "Hello championie Andy!" {
		t.Error("Incorrect user greeting")
	}
}
