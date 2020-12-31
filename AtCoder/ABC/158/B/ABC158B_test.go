package main

import (
	"testing"
)

func TestABC158B1(t *testing.T) {
	expect := uint64(4)
	actual := ABC158B(8, 3, 4)
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAB1058B2(t *testing.T) {
	expect := uint64(0)
	actual := ABC158B(8, 0, 4)
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC158B3(t *testing.T) {
	expect := uint64(2)
	actual := ABC158B(6, 2, 4)
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAB1058B4(t *testing.T) {
	expect := uint64(2)
	actual := ABC158B(8, 1, 4)
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
