package main

import (
	"testing"
)

func TestABC045C1(t *testing.T) {
	expect := 176
	actual := ABC045C("125")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC045C2(t *testing.T) {
	expect := 12656242944
	actual := ABC045C("9999999999")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
