package main

import (
	"testing"
)

func TestHitachi2020B1(t *testing.T) {
	expect := 5
	actual := Hitachi2020B([]int{3, 3}, []int{3, 3, 3}, []coupon{{1, 2, 1}})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestHitachi2020B2(t *testing.T) {
	expect := 10
	actual := Hitachi2020B([]int{10}, []int{10}, []coupon{{1, 1, 5}, {1, 1, 10}})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestHitachi2020B3(t *testing.T) {
	expect := 6
	actual := Hitachi2020B([]int{3, 5}, []int{3, 5}, []coupon{{2, 2, 2}})
	if expect != actual {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
