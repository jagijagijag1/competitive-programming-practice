package main

import (
	"testing"
)

func TestABC142C1(t *testing.T) {
	expect := []int{3, 1, 2}
	actual := ABC142C([]int{2, 3, 1})
	if !intSliceEq(actual, expect) {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC142C2(t *testing.T) {
	expect := []int{1, 2, 3, 4, 5}
	actual := ABC142C([]int{1, 2, 3, 4, 5})
	if !intSliceEq(actual, expect) {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC142C3(t *testing.T) {
	expect := []int{8, 2, 4, 5, 6, 7, 3, 1}
	actual := ABC142C([]int{8, 2, 7, 3, 4, 5, 6, 1})
	if !intSliceEq(actual, expect) {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC142C4(t *testing.T) {
	expect := []int{1}
	actual := ABC142C([]int{1})
	if !intSliceEq(actual, expect) {
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
