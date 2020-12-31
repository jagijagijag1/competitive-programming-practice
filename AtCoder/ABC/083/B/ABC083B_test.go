package main

import (
	"testing"
)

func TestABC083B1(t *testing.T) {
	expect := 84
	actual := ABC083B(20, 2, 5)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC083B2(t *testing.T) {
	expect := 13
	actual := ABC083B(10, 1, 2)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC083B3(t *testing.T) {
	expect := 4554
	actual := ABC083B(100, 4, 16)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
