package main

import (
	"testing"
)

func TestABC085B1(t *testing.T) {
	expect := 3
	actual := ABC085B([]int{10, 8, 8, 6})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC085B2(t *testing.T) {
	expect := 1
	actual := ABC085B([]int{15, 15, 15})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC085B3(t *testing.T) {
	expect := 4
	actual := ABC085B([]int{50, 30, 50, 100, 50, 80, 30})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
