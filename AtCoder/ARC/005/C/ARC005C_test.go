package main

import (
	"testing"
)

func TestARC005C1(t *testing.T) {
	expect := "YES"
	actual := ARC005C(4, 5, []string{
		"s####",
		"....#",
		"#####",
		"#...g",
	})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC005C2(t *testing.T) {
	expect := "YES"
	actual := ARC005C(4, 4, []string{
		"...s",
		"....",
		"....",
		".g..",
	})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC005C3(t *testing.T) {
	expect := "YES"
	actual := ARC005C(10, 10, []string{
		"s.........",
		"#########.",
		"#.......#.",
		"#..####.#.",
		"##....#.#.",
		"#####.#.#.",
		"g##.#.#.#.",
		"###.#.#.#.",
		"###.#.#.#.",
		"#.....#...",
	})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC005C4(t *testing.T) {
	expect := "YES"
	actual := ARC005C(6, 6, []string{
		".....s",
		"###...",
		"###...",
		"######",
		"...###",
		"g.####",
	})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC005C5(t *testing.T) {
	expect := "NO"
	actual := ARC005C(1, 10, []string{
		"s..####..g",
	})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
