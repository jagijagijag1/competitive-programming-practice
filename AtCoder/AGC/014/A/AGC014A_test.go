package main

import (
	"testing"
)

func TestAGC014A1(t *testing.T) {
	expect := 3
	actual := AGC014A(4, 12, 20)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC014A2(t *testing.T) {
	expect := -1
	actual := AGC014A(14, 14, 14)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC014A3(t *testing.T) {
	expect := 1
	actual := AGC014A(454, 414, 444)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC014A4(t *testing.T) {
	expect := 8
	actual := AGC014A(1000000000, 9000000000, 900000000)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
