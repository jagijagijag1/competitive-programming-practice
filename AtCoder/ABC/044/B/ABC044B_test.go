package main

import (
	"testing"
)

func TestABC044B1(t *testing.T) {
	expect := "Yes"
	actual := ABC044B("abaccaba")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC044B2(t *testing.T) {
	expect := "No"
	actual := ABC044B("hthth")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
