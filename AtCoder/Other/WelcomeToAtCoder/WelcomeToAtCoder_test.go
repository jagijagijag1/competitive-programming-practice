package main

import (
	"testing"
)

func TestWelcomeToAtCoder1(t *testing.T) {
	expect := "6 test"
	actual := WelcomeToAtCoder(1, 2, 3, "test")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestWelcomeToAtCoder2(t *testing.T) {
	expect := "456 myonmyon"
	actual := WelcomeToAtCoder(72, 128, 256, "myonmyon")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
