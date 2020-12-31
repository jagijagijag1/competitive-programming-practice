package main

import (
	"testing"
)

func TestABC133B1(t *testing.T) {
	expect := 1
	actual := ABC133B([][]int{{1, 2}, {5, 5}, {-2, 8}})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC133B2(t *testing.T) {
	expect := 2
	actual := ABC133B([][]int{{-3, 7, 8, 2}, {-12, 1, 10, 2}, {-2, 8, 9, 3}})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC133B3(t *testing.T) {
	expect := 10
	actual := ABC133B([][]int{{1}, {2}, {3}, {4}, {5}})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
