package main

import (
	"testing"
)

func TestABC070B1(t *testing.T) {
	expect := 50
	actual := ABC070B(0, 75, 25, 100)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC070B2(t *testing.T) {
	expect := 0
	actual := ABC070B(0, 33, 66, 99)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC070B3(t *testing.T) {
	expect := 60
	actual := ABC070B(10, 90, 20, 80)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC070B4(t *testing.T) {
	expect := 14
	actual := ABC070B(0, 14, 0, 28)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
