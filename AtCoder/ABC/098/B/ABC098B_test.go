package main

import (
	"testing"
)

func TestABC098B1(t *testing.T) {
	expect := 2
	actual := ABC098B("aabbca")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC098B2(t *testing.T) {
	expect := 1
	actual := ABC098B("aaaaaaaaaa")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC098B3(t *testing.T) {
	expect := 9
	actual := ABC098B("tgxgdqkyjzhyputjjtllptdfxocrylqfqjynmfbfucbir")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
