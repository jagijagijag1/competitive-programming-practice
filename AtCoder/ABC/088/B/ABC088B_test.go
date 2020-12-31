package main

import (
	"testing"
)

func TestABC088B1(t *testing.T) {
	expect := 2
	actual := ABC088B([]int{3, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC088B2(t *testing.T) {
	expect := 5
	actual := ABC088B([]int{2, 7, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC088B3(t *testing.T) {
	expect := 18
	actual := ABC088B([]int{20, 18, 2, 18})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
