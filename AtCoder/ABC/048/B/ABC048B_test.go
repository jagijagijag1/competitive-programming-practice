package main

import (
	"testing"
)

func TestABC048B1(t *testing.T) {
	expect := 3
	actual := ABC048B(4, 8, 2)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC048B2(t *testing.T) {
	expect := 6
	actual := ABC048B(0, 5, 1)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC048B3(t *testing.T) {
	expect := 333333333333333333
	actual := ABC048B(1, 1000000000000000000, 3)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC048B4(t *testing.T) {
	expect := 1
	actual := ABC048B(0, 0, 2)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC048B5(t *testing.T) {
	expect := 1
	actual := ABC048B(0, 1, 2)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC048B6(t *testing.T) {
	expect := 2
	actual := ABC048B(0, 1000000000000000000, 1000000000000000000)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
