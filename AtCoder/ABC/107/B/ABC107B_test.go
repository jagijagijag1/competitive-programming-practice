package main

import (
	"testing"
)

func TestABC107B1(t *testing.T) {
	expect := []string{"###", "###", ".##"}
	actual := ABC107B([]string{"##.#", "....", "##.#", ".#.#"})
	if !stringSliceEq(expect, actual) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC107B2(t *testing.T) {
	expect := []string{"#..", ".#.", "..#"}
	actual := ABC107B([]string{"#..", ".#.", "..#"})
	if !stringSliceEq(expect, actual) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC107B3(t *testing.T) {
	expect := []string{"#"}
	actual := ABC107B([]string{".....", ".....", "..#..", ".....", "....."})
	if !stringSliceEq(expect, actual) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC107B4(t *testing.T) {
	expect := []string{"..#", "#..", ".#.", ".#.", "#.#"}
	actual := ABC107B([]string{"......", "....#.", ".#....", "..#...", "..#...", "......", ".#..#."})
	if !stringSliceEq(expect, actual) {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC107B5(t *testing.T) {
	expect := []string{""}
	actual := ABC107B([]string{".....", ".....", ".....", ".....", "....."})
	if !stringSliceEq(expect, actual) {
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
