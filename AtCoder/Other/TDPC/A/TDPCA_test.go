package main

import (
	"testing"
)

func TestTDPCA1(t *testing.T) {
	expect := 7
	actual := TDPCA([]int{2, 3, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestTDPCA2(t *testing.T) {
	expect := 11
	actual := TDPCA([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
