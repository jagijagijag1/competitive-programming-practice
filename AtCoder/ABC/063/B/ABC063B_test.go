package main

import (
	"testing"
)

func TestABC063B1(t *testing.T) {
	expect := "yes"
	actual := ABC063B("uncopyrightable")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC063B2(t *testing.T) {
	expect := "no"
	actual := ABC063B("different")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC063B3(t *testing.T) {
	expect := "yes"
	actual := ABC063B("no")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
