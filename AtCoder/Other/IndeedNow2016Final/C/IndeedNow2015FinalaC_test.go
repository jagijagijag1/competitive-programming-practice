package main

import (
	"testing"
)

func TestIndeedNow2015FinalaC1(t *testing.T) {
	expect := []int{6, 8, 0, 3, 0, 8}
	actual := IndeedNow2015FinalaC(3, 6,
		[]IndeedNowCompany{{1, 2, 3, 3}, {3, 3, 3, 6}, {4, 4, 4, 8}},
		[]IndeedNowApplicant{{3, 4, 3}, {4, 4, 4}, {100, 100, 1}, {2, 3, 4}, {0, 0, 0}, {100, 100, 100}},
	)
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

func stringSliceEq(a, b []string) bool {
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
