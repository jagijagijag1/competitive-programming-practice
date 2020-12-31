package main

import (
	"testing"
)

func TestABC007C1(t *testing.T) {
	expect := 11
	actual := ABC007C(7, 8, 2, 2, 4, 5, []string{"########", "#......#", "#.######", "#..#...#", "#..##..#", "##.....#", "########"})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC007C2(t *testing.T) {
	expect := 10
	actual := ABC007C(5, 8, 2, 2, 2, 4, []string{"########", "#.#....#", "#.###..#", "#......#", "########"})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
