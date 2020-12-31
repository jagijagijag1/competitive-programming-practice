package main

import (
	"testing"
)

func TestABC149C1(t *testing.T) {
	expect := 23
	actual := ABC149C(20)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC149C2(t *testing.T) {
	expect := 2
	actual := ABC149C(2)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC149C3(t *testing.T) {
	expect := 100003
	actual := ABC149C(99992)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
