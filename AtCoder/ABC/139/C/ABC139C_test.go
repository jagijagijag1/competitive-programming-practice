package main

import (
	"testing"
)

func TestABC139C1(t *testing.T) {
	expect := 2
	actual := ABC139C([]int64{10, 4, 8, 7, 3})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC139C2(t *testing.T) {
	expect := 3
	actual := ABC139C([]int64{4, 4, 5, 6, 6, 5, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC139C3(t *testing.T) {
	expect := 0
	actual := ABC139C([]int64{1, 2, 3, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
