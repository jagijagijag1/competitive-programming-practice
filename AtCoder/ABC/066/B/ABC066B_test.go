package main

import (
	"testing"
)

func TestABC066B1(t *testing.T) {
	expect := 6
	actual := ABC066B("abaababaab")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC066B2(t *testing.T) {
	expect := 2
	actual := ABC066B("xxxx")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC066B3(t *testing.T) {
	expect := 6
	actual := ABC066B("abcabcabcabc")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC066B4(t *testing.T) {
	expect := 14
	actual := ABC066B("akasakaakasakasakaakas")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
