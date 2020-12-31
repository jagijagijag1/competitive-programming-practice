package main

import (
	"testing"
)

func TestABC038D1(t *testing.T) {
	expect := 3
	actual := ABC038D([]int{3, 1, 2}, []int{3, 1, 2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC038D2(t *testing.T) {
	expect := 1
	actual := ABC038D([]int{4, 4}, []int{5, 3})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC038D3(t *testing.T) {
	expect := 3
	actual := ABC038D([]int{2, 3, 4, 6}, []int{5, 3, 5, 6})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC038D4(t *testing.T) {
	expect := 4
	actual := ABC038D([]int{8, 5, 2, 4, 2}, []int{8, 3, 2, 2, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
