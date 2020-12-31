package main

import (
	"testing"
)

func TestABC139D1(t *testing.T) {
	expect := 1
	actual := ABC139D(2)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC139D2(t *testing.T) {
	expect := 78
	actual := ABC139D(13)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC139D3(t *testing.T) {
	expect := 0
	actual := ABC139D(1)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
