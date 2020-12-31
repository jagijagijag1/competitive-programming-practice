package main

import (
	"testing"
)

func TestABC127C1(t *testing.T) {
	expect := 2
	actual := ABC127C(4, []int{1, 2}, []int{3, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC127C2(t *testing.T) {
	expect := 1
	actual := ABC127C(10, []int{3, 5, 6}, []int{6, 7, 9})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC127C3(t *testing.T) {
	expect := 100000
	actual := ABC127C(100000, []int{1}, []int{100000})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC127C4(t *testing.T) {
	expect := 0
	actual := ABC127C(10, []int{3, 6}, []int{5, 7})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
