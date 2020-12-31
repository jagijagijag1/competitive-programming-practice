package main

import (
	"testing"
)

func TestABC055B1(t *testing.T) {
	expect := 6
	actual := ABC055B(3)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC055B2(t *testing.T) {
	expect := 3628800
	actual := ABC055B(10)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC055B3(t *testing.T) {
	expect := 457992974
	actual := ABC055B(100000)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
