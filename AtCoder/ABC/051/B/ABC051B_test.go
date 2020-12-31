package main

import (
	"testing"
)

func TestABC051B1(t *testing.T) {
	expect := 6
	actual := ABC051B(2, 2)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC051B2(t *testing.T) {
	expect := 1
	actual := ABC051B(5, 15)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
