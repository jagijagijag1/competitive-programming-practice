package main

import (
	"testing"
)

func TestABC138C1(t *testing.T) {
	expect := 3.5
	actual := ABC138C([]int{3, 4})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}

func TestABC138C2(t *testing.T) {
	expect := 375.0
	actual := ABC138C([]int{500, 300, 200})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}

func TestABC138C3(t *testing.T) {
	expect := 138.0
	actual := ABC138C([]int{138, 138, 138, 138, 138})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}
