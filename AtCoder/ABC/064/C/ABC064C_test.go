package main

import (
	"testing"
)

func TestABC064C1(t *testing.T) {
	expect1, expect2 := 2, 2
	actual1, actual2 := ABC064C([]int{2100, 2500, 2700, 2700})
	if actual1 != expect1 || actual2 != expect2 {
		t.Errorf(`expect="%d, %d" actual="%d, %d"`, expect1, expect2, actual1, actual2)
	}
}

func TestABC064C2(t *testing.T) {
	expect1, expect2 := 3, 5
	actual1, actual2 := ABC064C([]int{1100, 1900, 2800, 3200, 3200})
	if actual1 != expect1 || actual2 != expect2 {
		t.Errorf(`expect="%d, %d" actual="%d, %d"`, expect1, expect2, actual1, actual2)
	}
}

func TestABC064C3(t *testing.T) {
	expect1, expect2 := 1, 1
	actual1, actual2 := ABC064C([]int{800, 810, 820, 830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940, 950, 960, 970, 980, 990})
	if actual1 != expect1 || actual2 != expect2 {
		t.Errorf(`expect="%d, %d" actual="%d, %d"`, expect1, expect2, actual1, actual2)
	}
}
