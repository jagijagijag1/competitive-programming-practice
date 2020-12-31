package main

import (
	"testing"
)

func TestABC161C1(t *testing.T) {
	expect := int64(1)
	actual := ABC161C(7, 4)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC161C2(t *testing.T) {
	expect := int64(2)
	actual := ABC161C(2, 6)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC161C3(t *testing.T) {
	expect := int64(0)
	actual := ABC161C(1000000000000000000, 1)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC161C4(t *testing.T) {
	expect := int64(0)
	actual := ABC161C(999999999999999, 3)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
