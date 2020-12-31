package main

import (
	"testing"
)

func TestARC098D1(t *testing.T) {
	expect := 5
	actual := ARC098D([]int{2, 5, 4, 6})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC098D2(t *testing.T) {
	expect := 45
	actual := ARC098D([]int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC098D3(t *testing.T) {
	expect := 37
	actual := ARC098D([]int{885, 8, 1, 128, 83, 32, 256, 206, 639, 16, 4, 128, 689, 32, 8, 64, 885, 969, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
