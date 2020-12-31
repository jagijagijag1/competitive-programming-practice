package main

import (
	"testing"
)

func TestAGC027A1(t *testing.T) {
	expect := 2
	actual := AGC027A(70, []int{20, 30, 10})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC027A2(t *testing.T) {
	expect := 1
	actual := AGC027A(10, []int{20, 30, 10})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC027A3(t *testing.T) {
	expect := 4
	actual := AGC027A(1111, []int{1, 10, 100, 1000})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC027A4(t *testing.T) {
	expect := 0
	actual := AGC027A(10, []int{20, 20})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC027A5(t *testing.T) {
	expect := 2
	actual := AGC027A(40, []int{20, 20, 20})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
