package main

import (
	"testing"
)

func TestABC05C1(t *testing.T) {
	expect := "yes"
	actual := ABC05C(1, []int{1, 2, 3}, []int{2, 3, 4})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC05C2(t *testing.T) {
	expect := "no"
	actual := ABC05C(1, []int{1, 2, 3}, []int{2, 3, 5})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC05C3(t *testing.T) {
	expect := "no"
	actual := ABC05C(1, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC05C4(t *testing.T) {
	expect := "no"
	actual := ABC05C(1, []int{1, 2, 3}, []int{1, 2, 2})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC05C5(t *testing.T) {
	expect := "yes"
	actual := ABC05C(2, []int{1, 3, 6, 10, 15}, []int{4, 8, 16})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
