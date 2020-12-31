package main

import (
	"testing"
)

func TestARC029A1(t *testing.T) {
	expect := 14
	actual := ARC029A([]int{4, 6, 7, 10})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC029A2(t *testing.T) {
	expect := 4
	actual := ARC029A([]int{1, 2, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC029A3(t *testing.T) {
	expect := 29
	actual := ARC029A([]int{29})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
