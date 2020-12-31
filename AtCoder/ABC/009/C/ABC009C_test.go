package main

import (
	"testing"
)

func TestABC009C1(t *testing.T) {
	expect := "abc"
	actual := ABC009C(2, "abc")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC009C2(t *testing.T) {
	expect := "actoder"
	actual := ABC009C(2, "atcoder")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC009C3(t *testing.T) {
	expect := "acdeort"
	actual := ABC009C(7, "atcoder")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC009C4(t *testing.T) {
	expect := "dehloworll"
	actual := ABC009C(3, "helloworld")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
