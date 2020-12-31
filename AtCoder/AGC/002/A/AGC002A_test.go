package main

import (
	"testing"
)

func TestAGC002A1(t *testing.T) {
	expect := "Positive"
	actual := AGC002A(1, 3)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestAGC002A2(t *testing.T) {
	expect := "Negative"
	actual := AGC002A(-3, -1)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestAGC002A3(t *testing.T) {
	expect := "Zero"
	actual := AGC002A(-1, 1)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestAGC002A4(t *testing.T) {
	expect := "Positive"
	actual := AGC002A(-3, -2)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
