package main

import (
	"testing"
)

func TestABC148C1(t *testing.T) {
	expect := 6
	actual := ABC148C(2, 3)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC148C2(t *testing.T) {
	expect := 18696
	actual := ABC148C(123, 456)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC148C3(t *testing.T) {
	expect := 9999900000
	actual := ABC148C(100000, 99999)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
