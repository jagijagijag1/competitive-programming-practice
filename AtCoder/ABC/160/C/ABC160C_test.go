package main

import (
	"testing"
)

func TestABC160C1(t *testing.T) {
	expect := 10
	actual := ABC160C(20, []int{5, 10, 15})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC160C2(t *testing.T) {
	expect := 10
	actual := ABC160C(20, []int{0, 5, 15})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
