package main

import (
	"testing"
)

func TestABC128C1(t *testing.T) {
	expect := 1
	actual := ABC128C(2, 2, [][]int{{1, 2}, {2}}, []int{0, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC128C2(t *testing.T) {
	expect := 0
	actual := ABC128C(2, 3, [][]int{{1, 2}, {1}, {2}}, []int{0, 0, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC128C3(t *testing.T) {
	expect := 8
	actual := ABC128C(5, 2, [][]int{{1, 2, 5}, {2, 3}}, []int{1, 0})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
