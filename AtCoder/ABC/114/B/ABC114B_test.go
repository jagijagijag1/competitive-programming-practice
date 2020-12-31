package main

import (
	"testing"
)

func TestABC114B1(t *testing.T) {
	expect := 34
	actual := ABC114B("1234567876")
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC114B2(t *testing.T) {
	expect := 0
	actual := ABC114B("35753")
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC114B3(t *testing.T) {
	expect := 642
	actual := ABC114B("1111111111")
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
