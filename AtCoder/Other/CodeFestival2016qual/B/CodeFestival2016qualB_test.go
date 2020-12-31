package main

import (
	"testing"
)

func TestCodeFestival2016qualB1(t *testing.T) {
	expect := 2
	actual := CodeFestival2016qualB([]int{2, 1, 4, 3})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestCodeFestival2016qualB2(t *testing.T) {
	expect := 0
	actual := CodeFestival2016qualB([]int{2, 3, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestCodeFestival2016qualB3(t *testing.T) {
	expect := 1
	actual := CodeFestival2016qualB([]int{5, 5, 5, 5, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
