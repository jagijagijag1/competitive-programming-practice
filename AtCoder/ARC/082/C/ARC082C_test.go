package main

import (
	"testing"
)

func TestARC082C1(t *testing.T) {
	expect := 4
	actual := ARC082C([]int{3, 1, 4, 1, 5, 9, 2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC082C2(t *testing.T) {
	expect := 3
	actual := ARC082C([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC082C3(t *testing.T) {
	expect := 1
	actual := ARC082C([]int{99999})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
