package main

import (
	"testing"
)

func TestABC049C1(t *testing.T) {
	expect := "YES"
	actual := ABC049C("erasedream")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC049C2(t *testing.T) {
	expect := "YES"
	actual := ABC049C("dreameraser")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC049C3(t *testing.T) {
	expect := "NO"
	actual := ABC049C("dreamerer")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
