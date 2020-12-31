package main

import (
	"testing"
)

func TestABC114C1(t *testing.T) {
	expect := 4
	actual := ABC114C("575")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC114C2(t *testing.T) {
	expect := 13
	actual := ABC114C("3600")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC114C3(t *testing.T) {
	expect := 26484
	actual := ABC114C("999999999")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
