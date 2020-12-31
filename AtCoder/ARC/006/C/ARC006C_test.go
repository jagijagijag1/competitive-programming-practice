package main

import (
	"testing"
)

func TestARC006C1(t *testing.T) {
	expect := 2
	actual := ARC006C([]int{4, 3, 1, 2, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC006C2(t *testing.T) {
	expect := 3
	actual := ARC006C([]int{93, 249, 150, 958, 442, 391, 25})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC006C3(t *testing.T) {
	expect := 1
	actual := ARC006C([]int{100, 100, 100, 100})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC006C4(t *testing.T) {
	expect := 6
	actual := ARC006C([]int{5, 10, 15, 20, 25, 30})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC006C5(t *testing.T) {
	expect := 6
	actual := ARC006C([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
