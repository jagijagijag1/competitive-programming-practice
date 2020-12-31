package main

import (
	"testing"
)

func TestABC098C1(t *testing.T) {
	expect := 1
	actual := ABC098C("WEEWW")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC098C2(t *testing.T) {
	expect := 4
	actual := ABC098C("WEWEWEEEWWWE")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC098C3(t *testing.T) {
	expect := 3
	actual := ABC098C("WWWWWEEE")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
