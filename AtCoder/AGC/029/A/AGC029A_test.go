package main

import (
	"testing"
)

func TestAGC029A1(t *testing.T) {
	expect := 2
	actual := AGC029A("BBW")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC029A2(t *testing.T) {
	expect := 6
	actual := AGC029A("BWBWBW")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
