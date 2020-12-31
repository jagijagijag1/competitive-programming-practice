package main

import (
	"testing"
)

func TestABC060B1(t *testing.T) {
	expect := "YES"
	actual := ABC060B(7, 5, 1)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC060B2(t *testing.T) {
	expect := "NO"
	actual := ABC060B(2, 2, 1)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC060B3(t *testing.T) {
	expect := "YES"
	actual := ABC060B(1, 100, 97)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC060B4(t *testing.T) {
	expect := "YES"
	actual := ABC060B(40, 98, 58)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC060B5(t *testing.T) {
	expect := "NO"
	actual := ABC060B(77, 42, 36)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
