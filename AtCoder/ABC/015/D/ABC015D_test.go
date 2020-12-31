package main

import (
	"testing"
)

func TestABC015D1(t *testing.T) {
	expect := 140
	actual := ABC015D(10, 3, 2, []int{4, 3, 6}, []int{100, 40, 20})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC015D2(t *testing.T) {
	expect := 18
	actual := ABC015D(10, 5, 4, []int{9, 3, 3, 2, 4}, []int{10, 7, 1, 6, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC015D3(t *testing.T) {
	expect := 210
	actual := ABC015D(22, 5, 3, []int{5, 8, 3, 4, 6}, []int{40, 50, 60, 70, 80})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
