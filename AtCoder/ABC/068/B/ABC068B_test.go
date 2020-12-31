package main

import (
	"testing"
)

func TestABC068B1(t *testing.T) {
	expect := 4
	actual := ABC068B(7)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC068B2(t *testing.T) {
	expect := 32
	actual := ABC068B(32)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC068B3(t *testing.T) {
	expect := 1
	actual := ABC068B(1)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC068B4(t *testing.T) {
	expect := 64
	actual := ABC068B(100)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
