package main

import (
	"testing"
)

func TestABC116C1(t *testing.T) {
	expect := 2
	actual := ABC116C(4, []int{1, 2, 2, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC116C2(t *testing.T) {
	expect := 5
	actual := ABC116C(5, []int{3, 1, 2, 3, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC116C3(t *testing.T) {
	expect := 221
	actual := ABC116C(8, []int{4, 23, 75, 0, 23, 96, 50, 100})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC116C4(t *testing.T) {
	expect := 4
	actual := ABC116C(4, []int{0, 2, 0, 2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
