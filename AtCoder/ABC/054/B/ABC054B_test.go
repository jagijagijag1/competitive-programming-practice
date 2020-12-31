package main

import (
	"testing"
)

func TestABC054B1(t *testing.T) {
	expect := "Yes"
	actual := ABC054B([]string{"#.#", ".#.", "#.#"}, []string{"#.", ".#"})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC054B2(t *testing.T) {
	expect := "No"
	actual := ABC054B([]string{"....", ".....", "....", "...."}, []string{"#"})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC054B3(t *testing.T) {
	expect := "Yes"
	actual := ABC054B([]string{"."}, []string{"."})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
