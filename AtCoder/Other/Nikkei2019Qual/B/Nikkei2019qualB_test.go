package main

import (
	"testing"
)

func TestNikkei2019qualB1(t *testing.T) {
	expect := 2
	actual := Nikkei2019qualB([]int{0, 1, 1, 2})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestNikkei2019qualB2(t *testing.T) {
	expect := 0
	actual := Nikkei2019qualB([]int{1, 1, 1, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestNikkei2019qualB3(t *testing.T) {
	expect := 24
	actual := Nikkei2019qualB([]int{0, 3, 2, 1, 2, 2, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestNikkei2019qualB4(t *testing.T) {
	expect := 1
	actual := Nikkei2019qualB([]int{0, 1, 2, 3, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestNikkei2019qualB5(t *testing.T) {
	expect := 0
	actual := Nikkei2019qualB([]int{0, 3, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
