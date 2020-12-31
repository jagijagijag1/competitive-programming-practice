package main

import (
	"testing"
)

func TestABC104C1(t *testing.T) {
	expect := 3
	actual := ABC104C(700, []int{3, 5}, []int{500, 800})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC104C2(t *testing.T) {
	expect := 7
	actual := ABC104C(2000, []int{3, 5}, []int{500, 800})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC104C3(t *testing.T) {
	expect := 2
	actual := ABC104C(400, []int{3, 5}, []int{500, 800})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC104C4(t *testing.T) {
	expect := 66
	actual := ABC104C(25000, []int{20, 40, 50, 30, 1}, []int{1000, 1000, 1000, 1000, 1000})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
