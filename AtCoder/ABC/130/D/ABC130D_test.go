package main

import (
	"testing"
)

func TestABC130D1(t *testing.T) {
	expect := uint(2)
	actual := ABC130D(10, []int{6, 1, 2, 7})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC130D2(t *testing.T) {
	expect := uint(3)
	actual := ABC130D(5, []int{3, 3, 3})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC130D3(t *testing.T) {
	expect := uint(36)
	actual := ABC130D(53462, []int{103, 35322, 232, 342, 21099, 90000, 18843, 9010, 35221, 19352})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC130D4(t *testing.T) {
	expect := uint(0)
	actual := ABC130D(100, []int{6, 1, 2, 7})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
