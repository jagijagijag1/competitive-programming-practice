package main

import (
	"testing"
)

func TestJSC2019QualB1(t *testing.T) {
	expect := 3
	actual := JSC2019QualB(2, []int{2, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestJSC2019QualB2(t *testing.T) {
	expect := 0
	actual := JSC2019QualB(5, []int{1, 1, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestJSC2019QualB3(t *testing.T) {
	expect := 185297239
	actual := JSC2019QualB(998244353, []int{10, 9, 8, 7, 5, 6, 3, 4, 2, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
