package main

import (
	"testing"
)

func TestARC065A1(t *testing.T) {
	expect := "YES"
	actual := ARC065A("erasedream")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC065A2(t *testing.T) {
	expect := "YES"
	actual := ARC065A("dreameraser")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC065A3(t *testing.T) {
	expect := "NO"
	actual := ARC065A("dreamerer")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
