package main

import (
	"testing"
)

func TestABC105B1(t *testing.T) {
	expect := 1
	actual := ABC105B(105)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC105B2(t *testing.T) {
	expect := 0
	actual := ABC105B(7)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
