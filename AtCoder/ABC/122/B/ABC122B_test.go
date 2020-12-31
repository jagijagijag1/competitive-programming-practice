package main

import (
	"testing"
)

func TestABC122B1(t *testing.T) {
	expect := 3
	actual := ABC122B("ATCODER")
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC122B2(t *testing.T) {
	expect := 5
	actual := ABC122B("HATAGAYA")
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC122B3(t *testing.T) {
	expect := 0
	actual := ABC122B("SHINJUKU")
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC122B4(t *testing.T) {
	expect := 0
	actual := ABC122B("ASSA")
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
