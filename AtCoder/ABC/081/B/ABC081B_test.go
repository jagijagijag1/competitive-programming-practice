package main

import (
	"testing"
)

func TestABC081B1(t *testing.T) {
	expect := 2
	actual := ABC081B([]int{8, 12, 40})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC081B2(t *testing.T) {
	expect := 0
	actual := ABC081B([]int{5, 6, 8, 10})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC081B3(t *testing.T) {
	expect := 8
	actual := ABC081B([]int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
