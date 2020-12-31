package main

import (
	"testing"
)

func TestABC087B1(t *testing.T) {
	expect := 2
	actual := ABC087B(2, 2, 2, 100)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC087B2(t *testing.T) {
	expect := 0
	actual := ABC087B(5, 1, 0, 150)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC087B3(t *testing.T) {
	expect := 213
	actual := ABC087B(30, 40, 50, 6000)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
