package main

import (
	"testing"
)

func TestABC144B1(t *testing.T) {
	expect := "Yes"
	actual := ABC144B(10)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC144B2(t *testing.T) {
	expect := "No"
	actual := ABC144B(50)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC144B3(t *testing.T) {
	expect := "Yes"
	actual := ABC144B(81)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
