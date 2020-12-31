package main

import (
	"testing"
)

func TestABC117B1(t *testing.T) {
	expect := "Yes"
	actual := ABC117B([]int{3, 8, 5, 1})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC117B2(t *testing.T) {
	expect := "No"
	actual := ABC117B([]int{3, 8, 4, 1})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC117B3(t *testing.T) {
	expect := "No"
	actual := ABC117B([]int{1, 8, 10, 5, 8, 12, 34, 100, 11, 3})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
