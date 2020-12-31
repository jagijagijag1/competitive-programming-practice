package main

import (
	"testing"
)

func TestARC088C1(t *testing.T) {
	expect := 3
	actual := ARC088C(3, 20)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC088C2(t *testing.T) {
	expect := 3
	actual := ARC088C(25, 100)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC088C3(t *testing.T) {
	expect := 31
	actual := ARC088C(314159265, 358979323846264338)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
