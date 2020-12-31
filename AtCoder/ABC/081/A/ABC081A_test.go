package main

import (
	"testing"
)

func TestABC081A1(t *testing.T) {
	expect := 2
	actual := ABC081A("101")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC081A2(t *testing.T) {
	expect := 0
	actual := ABC081A("000")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
