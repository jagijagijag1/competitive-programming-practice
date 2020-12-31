package main

import (
	"testing"
)

func TestABC148E1(t *testing.T) {
	expect := uint64(1)
	actual := ABC148E(12)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC148E2(t *testing.T) {
	expect := uint64(0)
	actual := ABC148E(5)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC148E3(t *testing.T) {
	expect := uint64(124999999999999995)
	actual := ABC148E(1000000000000000000)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC148Enaive(t *testing.T) {
	expect := uint64(0)
	actual := ABC148Enaive(32)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
