package main

import (
	"testing"
)

func TestABC113C1(t *testing.T) {
	expect := "000001000002¥n000002000001¥n000001000001"
	actual := ABC113C(2, 3, []int{1, 2, 1}, []int{32, 63, 12})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC113C2(t *testing.T) {
	expect := "000002000001¥n000002000002¥n000002000003"
	actual := ABC113C(2, 3, []int{2, 2, 2}, []int{55, 77, 99})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
