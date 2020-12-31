package main

import (
	"testing"
)

func TestTDPCD1(t *testing.T) {
	expect := 0.416666667
	actual := TDPCD(2, 6)
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}

func TestTDPCD2(t *testing.T) {
	expect := 0.875000000
	actual := TDPCD(3, 2)
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}
