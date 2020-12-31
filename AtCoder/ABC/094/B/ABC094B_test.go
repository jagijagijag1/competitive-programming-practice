package main

import (
	"testing"
)

func TestABC094B1(t *testing.T) {
	expect := 1
	actual := ABC094B(3, []int{1, 2, 4})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC094B2(t *testing.T) {
	expect := 0
	actual := ABC094B(2, []int{4, 5, 6})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC094B3(t *testing.T) {
	expect := 3
	actual := ABC094B(5, []int{1, 2, 3, 4, 6, 8, 9})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC094B4(t *testing.T) {
	expect := 0
	actual := ABC094B(5, []int{1})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
