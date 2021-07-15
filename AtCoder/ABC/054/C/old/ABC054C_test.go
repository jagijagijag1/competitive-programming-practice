package main

import (
	"testing"
)

func TestABC054C1(t *testing.T) {
	expect := 2
	actual := ABC054C(3, []int{1, 1, 2}, []int{2, 3, 3})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC054C2(t *testing.T) {
	expect := 1
	actual := ABC054C(7, []int{1, 2, 3, 4, 4, 5, 6}, []int{3, 7, 4, 5, 6, 6, 7})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
