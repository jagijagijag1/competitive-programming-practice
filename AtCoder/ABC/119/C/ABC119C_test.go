package main

import (
	"testing"
)

func TestABC119C1(t *testing.T) {
	expect := 23
	actual := ABC119C(100, 90, 80, []int{98, 40, 30, 21, 80})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC119C2(t *testing.T) {
	expect := 0
	actual := ABC119C(100, 90, 80, []int{100, 100, 90, 90, 90, 80, 80, 80})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC119C3(t *testing.T) {
	expect := 243
	actual := ABC119C(1000, 800, 100, []int{300, 333, 400, 444, 500, 555, 600, 666})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
