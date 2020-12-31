package main

import (
	"testing"
)

func TestABC118B1(t *testing.T) {
	expect := 1
	actual := ABC118B(4, [][]int{{1, 3}, {1, 2, 3}, {3, 2}})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC118B2(t *testing.T) {
	expect := 0
	actual := ABC118B(5, [][]int{{2, 3, 4, 5}, {1, 3, 4, 5}, {1, 2, 4, 5}, {1, 2, 3, 5}, {1, 2, 3, 4}})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC118B3(t *testing.T) {
	expect := 3
	actual := ABC118B(30, [][]int{{5, 10, 30}})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
