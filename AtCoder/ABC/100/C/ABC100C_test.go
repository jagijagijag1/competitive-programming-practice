package main

import (
	"testing"
)

func TestABC100C1(t *testing.T) {
	expect := 3
	actual := ABC100C([]int64{5, 2, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC100C2(t *testing.T) {
	expect := 0
	actual := ABC100C([]int64{631, 577, 243, 199})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC100C3(t *testing.T) {
	expect := 39
	actual := ABC100C([]int64{184, 2126, 1721, 1800, 1024, 2528, 3360, 1945, 1280, 1776})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
