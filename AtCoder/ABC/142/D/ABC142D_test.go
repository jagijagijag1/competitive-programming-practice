package main

import (
	"testing"
)

func TestABC142D1(t *testing.T) {
	expect := 3
	actual := ABC142D(12, 18)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC142D2(t *testing.T) {
	expect := 4
	actual := ABC142D(420, 660)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC142D3(t *testing.T) {
	expect := 1
	actual := ABC142D(1, 2019)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC142D4(t *testing.T) {
	expect := 3
	actual := ABC142D(11, 121)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC142D5(t *testing.T) {
	expect := 2
	actual := ABC142D(6442450941, 4294967294)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC142D6(t *testing.T) {
	expect := 1
	actual := ABC142D(1, 1)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
