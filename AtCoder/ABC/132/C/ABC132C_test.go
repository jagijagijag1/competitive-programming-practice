package main

import (
	"testing"
)

func TestABC132C1(t *testing.T) {
	expect := 2
	actual := ABC132C([]int{9, 1, 4, 4, 6, 7})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC132C2(t *testing.T) {
	expect := 0
	actual := ABC132C([]int{9, 1, 14, 5, 5, 4, 4, 14})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC132C3(t *testing.T) {
	expect := 42685
	actual := ABC132C([]int{99592, 10342, 29105, 78532, 83018, 11639, 92015, 77204, 30914, 21912, 34519, 80835, 100000, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC132C4(t *testing.T) {
	expect := 3
	actual := ABC132C([]int{2, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
