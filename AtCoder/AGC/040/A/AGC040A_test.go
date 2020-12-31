package main

import (
	"testing"
)

func TestAGC040A1(t *testing.T) {
	expect := 3
	actual := AGC040A("<>>")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestAGC040A2(t *testing.T) {
	expect := 28
	actual := AGC040A("<>>><<><<<<<>>><")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
