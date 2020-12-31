package main

import (
	"testing"
)

func TestABC153E1(t *testing.T) {
	expect := 4
	actual := ABC153E(9, 3, []int{8, 4, 2}, []int{3, 2, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC153E2(t *testing.T) {
	expect := 100
	actual := ABC153E(100, 6, []int{1, 2, 3, 4, 5, 6}, []int{1, 3, 9, 27, 81, 243})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC153E3(t *testing.T) {
	expect := 139815
	actual := ABC153E(9999, 10, []int{540, 691, 700, 510, 415, 551, 587, 619, 588, 176}, []int{7550, 9680, 9790, 7150, 5818, 7712, 8227, 8671, 8228, 2461})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC153E4(t *testing.T) {
	expect := 100000000
	actual := ABC153E(1000, 1, []int{1}, []int{1000})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC153E5(t *testing.T) {
	expect := 99990000
	actual := ABC153E(10000, 2, []int{1, 1}, []int{9999, 10000})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
