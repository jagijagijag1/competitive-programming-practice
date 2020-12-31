package main

import (
	"testing"
)

func TestABC134C1(t *testing.T) {
	expect := []int{4, 3, 4}
	actual := ABC134C([]int{1, 4, 3})
	if !intSliceEq(expect, actual) {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC134C2(t *testing.T) {
	expect := []int{5, 5}
	actual := ABC134C([]int{5, 5})
	if !intSliceEq(expect, actual) {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
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
