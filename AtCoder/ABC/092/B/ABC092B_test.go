package main

import (
	"testing"
)

func TestABC092B1(t *testing.T) {
	expect := 8
	actual := ABC092B(7, 1, []int{2, 5, 10})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC092B2(t *testing.T) {
	expect := 29
	actual := ABC092B(8, 20, []int{1, 10})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC092B3(t *testing.T) {
	expect := 56
	actual := ABC092B(30, 44, []int{26, 18, 81, 18, 6})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
