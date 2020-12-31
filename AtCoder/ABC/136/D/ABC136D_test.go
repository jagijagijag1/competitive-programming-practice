package main

import (
	"testing"
)

func TestABC136D1(t *testing.T) {
	expect := "0 1 2 1 1"
	actual := ABC136D("RRLRL")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC136D2(t *testing.T) {
	expect := "0 3 3 0 0 0 1 1 0 2 2 0"
	actual := ABC136D("RRLLLLRLRRLL")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC136D3(t *testing.T) {
	expect := "0 0 3 2 0 2 1 0 0 0 4 4 0 0 0 0"
	actual := ABC136D("RRRLLRLLRRRLLLLL")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
