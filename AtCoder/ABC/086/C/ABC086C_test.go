package main

import (
	"testing"
)

func TestABC086C1(t *testing.T) {
	expect := "Yes"
	actual := ABC086C([]int{3, 6}, []int{1, 1}, []int{2, 1})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC086C2(t *testing.T) {
	expect := "No"
	actual := ABC086C([]int{2}, []int{100}, []int{100})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC086C3(t *testing.T) {
	expect := "No"
	actual := ABC086C([]int{5, 10}, []int{1, 1}, []int{1, 1})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
