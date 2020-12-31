package main

import (
	"testing"
)

func TestABC074B1(t *testing.T) {
	expect := 4
	actual := ABC074B(10, []int{2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC074B2(t *testing.T) {
	expect := 12
	actual := ABC074B(9, []int{3, 6})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC074B3(t *testing.T) {
	expect := 74
	actual := ABC074B(20, []int{11, 12, 9, 17, 12})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
