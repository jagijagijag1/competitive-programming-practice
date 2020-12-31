package main

import (
	"testing"
)

func TestABC150C1(t *testing.T) {
	expect := 3
	actual := ABC150C(3, []int{1, 3, 2}, []int{3, 1, 2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC150C2(t *testing.T) {
	expect := 17517
	actual := ABC150C(8, []int{7, 3, 5, 4, 2, 1, 6, 8}, []int{3, 8, 2, 5, 4, 6, 7, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC150C3(t *testing.T) {
	expect := 0
	actual := ABC150C(3, []int{1, 2, 3}, []int{1, 2, 3})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
