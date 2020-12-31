package main

import (
	"testing"
)

func TestABC079C1(t *testing.T) {
	expect := "1+2+2+2=7"
	// actual := ABC079C(1, 2, 2, 2)
	actual := ABC079Cbitsearch(1, 2, 2, 2)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC079C2(t *testing.T) {
	expect := "0-2+9+0=7"
	actual := ABC079C(0, 2, 9, 0)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC079C3(t *testing.T) {
	expect := "3+2+4-2=7"
	actual := ABC079C(3, 2, 4, 2)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
