package main

import (
	"testing"
)

func TestARC100C1(t *testing.T) {
	expect := 2
	actual := ARC100C([]int{2, 2, 3, 5, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC100C2(t *testing.T) {
	expect := 0
	actual := ARC100C([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC100C3(t *testing.T) {
	expect := 18
	actual := ARC100C([]int{6, 5, 4, 3, 2, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC100C4(t *testing.T) {
	expect := 6
	actual := ARC100C([]int{1, 1, 1, 1, 2, 3, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
