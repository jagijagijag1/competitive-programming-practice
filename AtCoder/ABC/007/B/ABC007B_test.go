package main

import (
	"testing"
)

func TestABC007B1(t *testing.T) {
	expect := "xy"
	actual := ABC007B("xyz")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC007B2(t *testing.T) {
	expect := "b"
	actual := ABC007B("c")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC007B3(t *testing.T) {
	expect := "-1"
	actual := ABC007B("a")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC007B4(t *testing.T) {
	expect := "aaaa"
	actual := ABC007B("aaaaa")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
