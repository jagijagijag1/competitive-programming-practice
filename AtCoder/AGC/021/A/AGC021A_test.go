package main

import (
	"testing"
)

func TestAGC021A1(t *testing.T) {
	expect := 18
	actual := AGC021A("100")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC021A2(t *testing.T) {
	expect := 35
	actual := AGC021A("9995")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC021A3(t *testing.T) {
	expect := 137
	actual := AGC021A("3141592653589793")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
