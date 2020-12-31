package main

import (
	"testing"
)

func TestAGC033A1(t *testing.T) {
	expect := 2
	actual := AGC033A(3, 3, []string{"...", ".#.", "..."})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC033A2(t *testing.T) {
	expect := 3
	actual := AGC033A(6, 6, []string{
		"..#..#",
		"......",
		"#..#..",
		"......",
		".#....",
		"....#.",
	})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
