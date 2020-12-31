package main

import (
	"testing"
)

func TestABC117C1(t *testing.T) {
	expect := 5
	actual := ABC117C(2, []int{10, 12, 1, 2, 14})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC117C2(t *testing.T) {
	expect := 19
	actual := ABC117C(3, []int{-10, -3, 0, 9, -100, 2, 17})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC117C3(t *testing.T) {
	expect := 0
	actual := ABC117C(100, []int{-1000000})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC117C4(t *testing.T) {
	expect := 13
	actual := ABC117C(1, []int{10, 12, 1, 2, 14})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC117C5(t *testing.T) {
	expect := 0
	actual := ABC117C(5, []int{10, 12, 1, 2, 14})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC117C6(t *testing.T) {
	expect := 0
	actual := ABC117C(100, []int{10, 12, 1, 2, 14})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
