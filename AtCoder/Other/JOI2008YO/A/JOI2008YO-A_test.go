package main

import (
	"testing"
)

func TestJOI2008YOA1(t *testing.T) {
	expect := 4
	actual := JOI2008YOA(380)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestJOI2008YOA2(t *testing.T) {
	expect := 15
	actual := JOI2008YOA(1)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
