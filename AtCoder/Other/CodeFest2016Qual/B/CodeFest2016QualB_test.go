package main

import (
	"testing"
)

func TestCodeFest2016QualB1(t *testing.T) {
	expect := []string{"Yes", "Yes", "No", "No", "Yes", "Yes", "Yes", "No", "No", "No"}
	actual := CodeFest2016QualB(2, 3, "abccabaabb")
	if !stringSliceEq(expect, actual) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestCodeFest2016QualB2(t *testing.T) {
	expect := []string{"No", "Yes", "Yes", "Yes", "Yes", "No", "Yes", "Yes", "No", "Yes", "No", "No"}
	actual := CodeFest2016QualB(5, 2, "cabbabaacaba")
	if !stringSliceEq(expect, actual) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestCodeFest2016QualB3(t *testing.T) {
	expect := []string{"No", "No", "No", "No", "No"}
	actual := CodeFest2016QualB(2, 2, "ccccc")
	if !stringSliceEq(expect, actual) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestCodeFest2016QualB4(t *testing.T) {
	expect := []string{"Yes", "Yes", "Yes", "Yes", "No", "No"}
	actual := CodeFest2016QualB(2, 2, "baaaab")
	if !stringSliceEq(expect, actual) {
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
