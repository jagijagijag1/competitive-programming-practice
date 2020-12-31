package main

import (
	"testing"
)

func TestAGC043A1(t *testing.T) {
	expect := 1
	actual := AGC043A(3, 3, []string{".##", ".#.", "##."})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC043A2(t *testing.T) {
	expect := 2
	actual := AGC043A(2, 2, []string{"#.", ".#"})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC043A3(t *testing.T) {
	expect := 4
	actual := AGC043A(5, 5, []string{".#.#.", "#.#.#", ".#.#.", "#.#.#", ".#.#."})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
