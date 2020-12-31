package main

import (
	"testing"
)

func TestABC156C1(t *testing.T) {
	expect := 5
	actual := ABC156C(2, []int{1, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC156C2(t *testing.T) {
	expect := 2354
	actual := ABC156C(7, []int{14, 14, 2, 13, 56, 2, 37})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC156C3(t *testing.T) {
	expect := 0
	actual := ABC156C(7, []int{2, 2, 2, 2, 2, 2, 2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC156C4(t *testing.T) {
	expect := 0
	actual := ABC156C(1, []int{1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
