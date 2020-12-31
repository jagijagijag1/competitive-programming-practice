package main

import (
	"testing"
)

func TestABC086B1(t *testing.T) {
	expect := "Yes"
	actual := ABC086B(121)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC086B2(t *testing.T) {
	expect := "No"
	actual := ABC086B(100100)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC086B3(t *testing.T) {
	expect := "No"
	actual := ABC086B(1210)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC086B4(t *testing.T) {
	expect := "Yes"
	actual := ABC086B(9801)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
