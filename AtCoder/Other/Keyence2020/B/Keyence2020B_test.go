package main

import (
	"testing"
)

func TestKeyence2020B1(t *testing.T) {
	expect := 3
	actual := Keyence2020B([]int64{2, 4, 9, 100}, []int64{4, 3, 3, 5})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestKeyence2020B2(t *testing.T) {
	expect := 1
	actual := Keyence2020B([]int64{8, 1}, []int64{20, 10})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestKeyence2020B3(t *testing.T) {
	expect := 5
	actual := Keyence2020B([]int64{10, 2, 4, 6, 8}, []int64{1, 1, 1, 1, 1})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
