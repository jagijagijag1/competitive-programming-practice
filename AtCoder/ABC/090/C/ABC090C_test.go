package main

import (
	"testing"
)

func TestABC090C1(t *testing.T) {
	expect := 14
	actual := ABC090C([][]int{{3, 2, 2, 4, 1}, {1, 2, 2, 2, 1}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC090C2(t *testing.T) {
	expect := 5
	actual := ABC090C([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC090C3(t *testing.T) {
	expect := 29
	actual := ABC090C([][]int{{3, 3, 4, 5, 4, 5, 3}, {5, 3, 4, 4, 2, 3, 2}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC090C4(t *testing.T) {
	expect := 5
	actual := ABC090C([][]int{{2}, {3}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
