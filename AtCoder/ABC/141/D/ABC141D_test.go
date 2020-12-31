package main

import (
	"testing"
)

func TestABC141D1(t *testing.T) {
	expect := 9
	actual := ABC141D(3, []int{2, 13, 8})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC141D2(t *testing.T) {
	expect := 6
	actual := ABC141D(4, []int{1, 9, 3, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC141D3(t *testing.T) {
	expect := 0
	actual := ABC141D(100000, []int{1000000000})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC141D4(t *testing.T) {
	expect := 9500000000
	actual := ABC141D(1, []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
