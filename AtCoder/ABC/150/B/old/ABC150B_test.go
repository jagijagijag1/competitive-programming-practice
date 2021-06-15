package main

import (
	"testing"
)

func TestABC150B1(t *testing.T) {
	expect := 2
	actual := ABC150B("ZABCDBABCQ")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC150B2(t *testing.T) {
	expect := 0
	actual := ABC150B("THREEONEFOURONEFIVE")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC150B3(t *testing.T) {
	expect := 5
	actual := ABC150B("ABCCABCBABCCABACBCBBABCBCBCBCABCB")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
