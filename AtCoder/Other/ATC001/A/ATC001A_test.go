package main

import (
	"testing"
)

func TestATC001A1(t *testing.T) {
	expect := "No"
	actual := ATC001A(4, 5, []string{"s####", "....#", "#####", "#...g"})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestATC001A2(t *testing.T) {
	expect := "Yes"
	actual := ATC001A(4, 4, []string{"...s", "....", "....", ".g.."})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestATC001A3(t *testing.T) {
	expect := "No"
	actual := ATC001A(10, 10, []string{"s.........", "#########.", "#.......#.", "#..####.#.", "##....#.#.", "#####.#.#.", "g.#.#.#.#.", "#.#.#.#.#.", "###.#.#.#.", "#.....#..."})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestATC001A4(t *testing.T) {
	expect := "Yes"
	actual := ATC001A(10, 10, []string{"s.........", "#########.", "#.......#.", "#..####.#.", "##....#.#.", "#####.#.#.", "g.#.#.#.#.", "#.#.#.#.#.", "#.#.#.#.#.", "#.....#..."})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestATC001A5(t *testing.T) {
	expect := "No"
	actual := ATC001A(1, 10, []string{"s..####..g"})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
