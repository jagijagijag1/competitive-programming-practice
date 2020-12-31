package main

import (
	"testing"
)

func TestABC002D1(t *testing.T) {
	expect := 3
	actual := ABC002D(5, []int{1, 2, 1}, []int{2, 3, 3})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC002D2(t *testing.T) {
	expect := 2
	actual := ABC002D(5, []int{1, 2, 3}, []int{2, 3, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC002D3(t *testing.T) {
	expect := 4
	actual := ABC002D(7, []int{1, 1, 2, 4, 4, 4, 5, 5, 6}, []int{2, 3, 3, 5, 6, 7, 6, 7, 7})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC002D4(t *testing.T) {
	expect := 1
	actual := ABC002D(12, []int{}, []int{})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
