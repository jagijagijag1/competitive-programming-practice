package main

import (
	"testing"
)

func TestABC096C1(t *testing.T) {
	expect := "Yes"
	actual := ABC096C(3, 3, []string{".#.", "###", ".#."})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC096C2(t *testing.T) {
	expect := "No"
	actual := ABC096C(5, 5, []string{"#.#.#", ".#.#.", "#.#.#", ".#.#.", "#.#.#"})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
