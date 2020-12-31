package main

import (
	"testing"
)

func TestABC157B1(t *testing.T) {
	expect := "Yes"
	actual := ABC157B([][]int{{84, 97, 66}, {79, 89, 11}, {61, 59, 7}}, []int{89, 7, 87, 79, 24, 84, 30})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC157B2(t *testing.T) {
	expect := "No"
	actual := ABC157B([][]int{{41, 7, 46}, {26, 89, 2}, {78, 92, 8}}, []int{6, 45, 16, 57, 17})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC157B3(t *testing.T) {
	expect := "Yes"
	actual := ABC157B([][]int{{60, 88, 34}, {92, 41, 43}, {65, 73, 48}}, []int{60, 43, 88, 11, 48, 73, 65, 41, 92, 34})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
