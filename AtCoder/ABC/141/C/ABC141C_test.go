package main

import (
	"testing"
)

func TestABC141C1(t *testing.T) {
	expect := []string{"No", "No", "Yes", "No", "No", "No"}
	actual := ABC141C(6, 3, []int{3, 1, 3, 2})
	if !stringSliceEq(actual, expect) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC141C2(t *testing.T) {
	expect := []string{"Yes", "Yes", "Yes", "Yes", "Yes", "Yes"}
	actual := ABC141C(6, 5, []int{3, 1, 3, 2})
	if !stringSliceEq(actual, expect) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC141C3(t *testing.T) {
	expect := []string{"No", "No", "No", "No", "Yes", "No", "No", "No", "Yes", "No"}
	actual := ABC141C(10, 13, []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9})
	if !stringSliceEq(actual, expect) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
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
