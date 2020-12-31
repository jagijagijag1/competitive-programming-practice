package main

import (
	"testing"
)

func TestJOI2008HOC1(t *testing.T) {
	expect := uint(48)
	actual := JOI2008HOC(50, []int{3, 14, 15, 9})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestJOI2008HOC2(t *testing.T) {
	expect := uint(20)
	actual := JOI2008HOC(21, []int{16, 11, 2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
