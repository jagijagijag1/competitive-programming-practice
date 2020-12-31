package main

import (
	"testing"
)

func TestABC076C1(t *testing.T) {
	expect := "atcoder"
	actual := ABC076C("?tc????", "coder")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC076C2(t *testing.T) {
	expect := "UNRESTORABLE"
	actual := ABC076C("??p??d??", "abc")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC076C3(t *testing.T) {
	expect := "aaamxxx"
	actual := ABC076C("???m???", "xxx")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC076C4(t *testing.T) {
	expect := "UNRESTORABLE"
	actual := ABC076C("???", "abcde")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
