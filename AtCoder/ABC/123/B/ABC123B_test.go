package main

import (
	"testing"
)

func TestABC123B1(t *testing.T) {
	expect := 215
	actual := ABC123B(29, 20, 7, 35, 120)
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC123B2(t *testing.T) {
	expect := 481
	actual := ABC123B(101, 86, 119, 108, 57)
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC123B3(t *testing.T) {
	expect := 643
	actual := ABC123B(123, 123, 123, 123, 123)
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
