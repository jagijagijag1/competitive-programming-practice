package main

import (
	"testing"
)

func TestJOI2011YOE1(t *testing.T) {
	expect := 4
	actual := JOI2011YOE(3, 3, 1, []string{"S..", "...", "..1"})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestJOI2011YOE2(t *testing.T) {
	expect := 12
	actual := JOI2011YOE(4, 5, 2, []string{".X..1", "....X", ".XX.S", ".2.X."})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestJOI2011YOE3(t *testing.T) {
	expect := 91
	actual := JOI2011YOE(10, 10, 9, []string{".X...X.S.X", "6..5X..X1X", "...XXXX..X", "X..9X...X.", "8.X2X..X3X", "...XX.X4..", "XX....7X..", "X..X..XX..", "X...X.XX..", "..X......."})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
