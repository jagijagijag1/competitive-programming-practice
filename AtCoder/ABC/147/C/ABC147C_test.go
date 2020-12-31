package main

import (
	"testing"
)

func TestABC147C1(t *testing.T) {
	expect := 2
	actual := ABC147C(3, [][]int{{2}, {1}, {2}}, [][]int{{1}, {1}, {0}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC147C2(t *testing.T) {
	expect := 0
	actual := ABC147C(3, [][]int{{2, 3}, {3, 1}, {1, 2}}, [][]int{{1, 0}, {1, 0}, {1, 0}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC147C3(t *testing.T) {
	expect := 1
	actual := ABC147C(2, [][]int{{2}, {1}}, [][]int{{0}, {0}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
