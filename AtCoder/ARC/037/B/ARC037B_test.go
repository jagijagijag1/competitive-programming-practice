package main

import (
	"testing"
)

func TestARC037B1(t *testing.T) {
	expect := 1
	actual := ARC037B(8, []int{1, 2, 2, 5, 6, 6, 7}, []int{2, 3, 4, 6, 7, 8, 8})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC037B2(t *testing.T) {
	expect := 4
	actual := ARC037B(5, []int{1}, []int{2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC037B3(t *testing.T) {
	expect := 0
	actual := ARC037B(11, []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 11})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
