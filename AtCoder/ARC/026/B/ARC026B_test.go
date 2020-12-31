package main

import (
	"testing"
)

func TestARC026B1(t *testing.T) {
	expect := "Perfect"
	actual := ARC026B(6)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC026B2(t *testing.T) {
	expect := "Abundant"
	actual := ARC026B(24)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC026B3(t *testing.T) {
	expect := "Deficient"
	actual := ARC026B(27)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC026B4(t *testing.T) {
	expect := "Abundant"
	actual := ARC026B(945)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC026B5(t *testing.T) {
	expect := "Abundant"
	actual := ARC026B(100000)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC026B6(t *testing.T) {
	expect := "Abundant"
	actual := ARC026B(10000000000)
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
