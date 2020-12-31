package main

import (
	"testing"
)

func TestJOI2013YOD1(t *testing.T) {
	expect := 80
	actual := JOI2013YOD(3, 4, []int{31, 27, 35}, []int{20, 23, 21, 28}, []int{25, 29, 35, 33}, []int{30, 90, 60, 40})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestJOI2013YOD2(t *testing.T) {
	expect := 300
	actual := JOI2013YOD(5, 2, []int{26, 28, 32, 29, 34}, []int{30, 25}, []int{35, 30}, []int{0, 100})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
