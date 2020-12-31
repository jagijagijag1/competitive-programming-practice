package main

import (
	"testing"
)

func TestPanasonic2020D1(t *testing.T) {
	expect := []string{"a"}
	actual := Panasonic2020D(1)
	if !stringSliceEq(actual, expect) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestPanasonic2020D2(t *testing.T) {
	expect := []string{"aa", "ab"}
	actual := Panasonic2020D(2)
	if !stringSliceEq(actual, expect) {
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
