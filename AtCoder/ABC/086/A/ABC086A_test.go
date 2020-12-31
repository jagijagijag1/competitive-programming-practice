package main

import (
	"testing"
)

func TestABC086A1(t *testing.T) {
	expect := "Even"
	actual := ABC086A(3, 4)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC086A2(t *testing.T) {
	expect := "Odd"
	actual := ABC086A(1, 21)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
