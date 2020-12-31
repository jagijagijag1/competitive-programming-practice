package main

import (
	"testing"
)

func TestABC057B1(t *testing.T) {
	expect := []int{2, 1}
	actual := ABC057B([]int{2, 0}, []int{0, 0}, []int{-1, 1}, []int{0, 0})
	if !intSliceEq(expect, actual) {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC057B2(t *testing.T) {
	expect := []int{3, 1, 2}
	actual := ABC057B([]int{10, -10, 3}, []int{10, -10, 3}, []int{1, 2, 3, 3}, []int{2, 3, 5, 5})
	if !intSliceEq(expect, actual) {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC057B3(t *testing.T) {
	expect := []int{5, 4, 3, 2, 1}
	actual := ABC057B(
		[]int{-10000000, -10000000, 10000000, 10000000, 0},
		[]int{-10000000, 10000000, -10000000, 10000000, 0},
		[]int{0, 10000000, 10000000, -10000000, -10000000},
		[]int{0, 10000000, -10000000, 10000000, -10000000},
	)
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
