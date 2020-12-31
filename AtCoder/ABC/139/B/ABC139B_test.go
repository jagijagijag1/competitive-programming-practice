package main

import (
	"testing"
)

func TestABC139B1(t *testing.T) {
	expect := 3
	actual := ABC139B(4, 10)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC139B2(t *testing.T) {
	expect := 2
	actual := ABC139B(8, 9)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC139B3(t *testing.T) {
	expect := 1
	actual := ABC139B(8, 8)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
