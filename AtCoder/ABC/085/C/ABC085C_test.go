package main

import (
	"testing"
)

func TestABC085C1(t *testing.T) {
	expect := "4 0 5"
	actual := ABC085C(9, 45000)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC085C2(t *testing.T) {
	expect := "-1 -1 -1"
	actual := ABC085C(20, 196000)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC085C3(t *testing.T) {
	expect := "14 27 959"
	actual := ABC085C(1000, 1234000)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC085C4(t *testing.T) {
	expect := "2000 0 0"
	actual := ABC085C(2000, 20000000)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func intSliceEq(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
