package main

import (
	"testing"
)

func TestABC046B1(t *testing.T) {
	expect := 2
	actual := ABC046B(2, 2)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC046B2(t *testing.T) {
	expect := 10
	actual := ABC046B(1, 10)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC046B3(t *testing.T) {
	expect := 322828856
	actual := ABC046B(10, 8)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
