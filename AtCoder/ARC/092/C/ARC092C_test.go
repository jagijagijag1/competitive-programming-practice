package main

import (
	"testing"
)

func TestARC092C1(t *testing.T) {
	expect := 2
	// actual := ARC092C([]int{2, 3, 1}, []int{0, 1, 3}, []int{4, 0, 5}, []int{2, 4, 5})
	actual := ARC092Cdinic([]int{2, 3, 1}, []int{0, 1, 3}, []int{4, 0, 5}, []int{2, 4, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC092C2(t *testing.T) {
	expect := 2
	// actual := ARC092C([]int{0, 1, 5}, []int{0, 1, 2}, []int{2, 3, 4}, []int{3, 4, 5})
	actual := ARC092Cdinic([]int{0, 1, 5}, []int{0, 1, 2}, []int{2, 3, 4}, []int{3, 4, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC092C3(t *testing.T) {
	expect := 0
	// actual := ARC092C([]int{2, 3}, []int{2, 3}, []int{0, 1}, []int{0, 1})
	actual := ARC092Cdinic([]int{2, 3}, []int{2, 3}, []int{0, 1}, []int{0, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC092C4(t *testing.T) {
	expect := 5
	// actual := ARC092C([]int{0, 7, 2, 4, 1}, []int{0, 3, 2, 8, 6}, []int{8, 6, 5, 9, 3}, []int{5, 9, 4, 1, 7})
	actual := ARC092Cdinic([]int{0, 7, 2, 4, 1}, []int{0, 3, 2, 8, 6}, []int{8, 6, 5, 9, 3}, []int{5, 9, 4, 1, 7})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC092C5(t *testing.T) {
	expect := 4
	// actual := ARC092C([]int{0, 1, 5, 6, 7}, []int{0, 1, 5, 6, 7}, []int{2, 3, 4, 8, 9}, []int{2, 3, 4, 8, 9})
	actual := ARC092Cdinic([]int{0, 1, 5, 6, 7}, []int{0, 1, 5, 6, 7}, []int{2, 3, 4, 8, 9}, []int{2, 3, 4, 8, 9})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
