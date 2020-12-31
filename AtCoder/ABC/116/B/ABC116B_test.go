package main

import (
	"testing"
)

func TestABC116B1(t *testing.T) {
	expect := 5
	actual := ABC116B(8)
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC116B2(t *testing.T) {
	expect := 18
	actual := ABC116B(7)
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC116B3(t *testing.T) {
	expect := 114
	actual := ABC116B(54)
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
