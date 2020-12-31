package main

import (
	"testing"
)

func TestJOI2010YOD1(t *testing.T) {
	expect := 7
	actual := JOI2010YOD(2, []string{"1", "2", "12", "1"})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestJOI2010YOD2(t *testing.T) {
	expect := 68
	actual := JOI2010YOD(3, []string{"72", "2", "12", "7", "2", "1"})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
