package main

import (
	"testing"
)

func TestABC161D1(t *testing.T) {
	expect := uint64(23)
	actual := ABC161D(15)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC161D2(t *testing.T) {
	expect := uint64(1)
	actual := ABC161D(1)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC161D3(t *testing.T) {
	expect := uint64(21)
	actual := ABC161D(13)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC161D4(t *testing.T) {
	expect := uint64(3234566667)
	actual := ABC161D(100000)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
