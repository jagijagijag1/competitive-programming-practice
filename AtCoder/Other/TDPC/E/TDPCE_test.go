package main

import (
	"testing"
)

func TestTDPCE1(t *testing.T) {
	expect := 33
	actual := TDPCE(3, []int{1, 0, 0})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestTDPCE2(t *testing.T) {
	expect := 468357804
	actual := TDPCE(7, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestTDPCE3(t *testing.T) {
	expect := 211
	actual := TDPCE(3, []int{6, 3, 4})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
