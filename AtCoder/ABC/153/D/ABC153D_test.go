package main

import (
	"testing"
)

func TestABC153D1(t *testing.T) {
	expect := uint64(3)
	actual := ABC153D(2)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC153D2(t *testing.T) {
	expect := uint64(7)
	actual := ABC153D(4)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC153D3(t *testing.T) {
	expect := uint64(1099511627775)
	actual := ABC153D(1000000000000)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
