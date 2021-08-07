package main

import (
	"testing"
)

func TestABC152C1(t *testing.T) {
	expect := 3
	actual := ABC152C([]int{4, 2, 5, 1, 3})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC152C2(t *testing.T) {
	expect := 4
	actual := ABC152C([]int{4, 3, 2, 1})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC152C3(t *testing.T) {
	expect := 1
	actual := ABC152C([]int{1, 2, 3, 4, 5, 6})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC152C4(t *testing.T) {
	expect := 4
	actual := ABC152C([]int{5, 7, 4, 2, 6, 8, 1, 3})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC152C5(t *testing.T) {
	expect := 1
	actual := ABC152C([]int{1})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
