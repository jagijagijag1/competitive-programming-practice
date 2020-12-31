package main

import (
	"testing"
)

func TestABC106D1(t *testing.T) {
	expect := []int{3}
	actual := ABC106D(2, []int{1, 1, 2}, []int{1, 2, 2}, []int{1}, []int{2})
	if !intSliceEq(actual, expect) {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC106D2(t *testing.T) {
	expect := []int{1, 1}
	actual := ABC106D(10, []int{1, 2, 7}, []int{5, 8, 10}, []int{1, 3}, []int{7, 10})
	if !intSliceEq(actual, expect) {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC106D3(t *testing.T) {
	expect := []int{7, 9, 10, 6, 8, 9, 6, 7, 8, 10}
	actual := ABC106D(10, []int{1, 2, 4, 4, 4, 5, 6, 6, 7, 10}, []int{6, 9, 5, 7, 7, 8, 6, 7, 9, 10}, []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 1}, []int{8, 9, 10, 8, 9, 10, 8, 9, 10, 10})
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
