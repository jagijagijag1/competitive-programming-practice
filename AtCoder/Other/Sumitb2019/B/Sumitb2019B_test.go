package main

import (
	"testing"
)

func TestSumitb2019B1(t *testing.T) {
	expect := "400"
	actual := Sumitb2019B(432)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestSumitb2019B2(t *testing.T) {
	expect := ":("
	actual := Sumitb2019B(1079)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestSumitb2019B3(t *testing.T) {
	expect := "927"
	actual := Sumitb2019B(1001)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
