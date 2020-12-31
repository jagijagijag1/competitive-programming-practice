package main

import (
	"testing"
)

func TestARC004A1(t *testing.T) {
	expect := float64(3.605551)
	actual := ARC004A([]int{1, 2, 4}, []int{1, 4, 3})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}

func TestARC004A2(t *testing.T) {
	expect := 10.630146
	actual := ARC004A([]int{1, 4, 3, 2, 5, 9, 6, 0, 8, 7}, []int{8, 0, 7, 4, 9, 1, 2, 2, 6, 8})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}

func TestARC004A3(t *testing.T) {
	expect := 141.421356
	actual := ARC004A([]int{0, 0, 100, 100}, []int{0, 100, 0, 100})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}

func TestARC004A4(t *testing.T) {
	expect := 4.000000
	actual := ARC004A([]int{3, 1, 0, 4, 2}, []int{0, 0, 0, 0, 0})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}

func TestARC004A5(t *testing.T) {
	expect := 4.242641
	actual := ARC004A([]int{2, 0, 1, 3}, []int{2, 0, 1, 3})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}
