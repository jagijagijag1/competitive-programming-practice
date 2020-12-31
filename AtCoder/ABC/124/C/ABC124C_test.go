package main

import (
	"testing"
)

func TestABC124C1(t *testing.T) {
	expect := 1
	actual := ABC124C("000")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC124C2(t *testing.T) {
	expect := 3
	actual := ABC124C("10010010")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC124C3(t *testing.T) {
	expect := 0
	actual := ABC124C("0")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
