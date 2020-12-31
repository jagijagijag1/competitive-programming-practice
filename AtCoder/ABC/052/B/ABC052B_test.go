package main

import (
	"testing"
)

func TestABC052B1(t *testing.T) {
	expect := 2
	actual := ABC052B("IIDID")
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC052B2(t *testing.T) {
	expect := 0
	actual := ABC052B("DDIDDII")
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
