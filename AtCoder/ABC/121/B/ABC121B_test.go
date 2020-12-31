package main

import (
	"testing"
)

func TestABC121B1(t *testing.T) {
	expect := 1
	actual := ABC121B([][]int{{3, 2, 1}, {1, 2, 2}}, []int{1, 2, 3}, -10)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC121B2(t *testing.T) {
	expect := 2
	actual := ABC121B([][]int{{100, 41}, {100, 40}, {-3, 0}, {-6, -2}, {18, -13}}, []int{-2, 5}, -4)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC121B3(t *testing.T) {
	expect := 0
	actual := ABC121B([][]int{{0, 100, 100}, {100, 100, 100}, {-100, 100, 100}}, []int{100, -100, 0}, 0)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
