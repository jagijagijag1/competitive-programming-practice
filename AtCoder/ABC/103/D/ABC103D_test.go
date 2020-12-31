package main

import (
	"testing"
)

func TestABC103D1(t *testing.T) {
	expect := 1
	actual := ABC103D(5, []int{1, 2}, []int{4, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC103D2(t *testing.T) {
	expect := 2
	actual := ABC103D(9, []int{1, 2, 3, 4, 7}, []int{8, 7, 5, 6, 9})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC103D3(t *testing.T) {
	expect := 4
	actual := ABC103D(5, []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4}, []int{2, 3, 4, 5, 3, 4, 5, 4, 5, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
