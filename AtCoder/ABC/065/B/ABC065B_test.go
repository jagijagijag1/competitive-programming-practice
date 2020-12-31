package main

import (
	"testing"
)

func TestABC065B1(t *testing.T) {
	expect := 2
	actual := ABC065B([]int{3, 1, 2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC065B2(t *testing.T) {
	expect := -1
	actual := ABC065B([]int{3, 4, 1, 2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC065B3(t *testing.T) {
	expect := 3
	actual := ABC065B([]int{3, 3, 4, 2, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC065B4(t *testing.T) {
	expect := -1
	actual := ABC065B([]int{3, 4, 4, 3})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
