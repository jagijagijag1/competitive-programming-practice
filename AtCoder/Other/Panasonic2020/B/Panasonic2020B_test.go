package main

import (
	"testing"
)

func TestPanasonic2020B1(t *testing.T) {
	expect := int64(10)
	actual := Panasonic2020B(4, 5)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestPanasonic2020B2(t *testing.T) {
	expect := int64(11)
	actual := Panasonic2020B(7, 3)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestPanasonic2020B3(t *testing.T) {
	expect := int64(500000000000000000)
	actual := Panasonic2020B(1000000000, 1000000000)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestPanasonic2020B4(t *testing.T) {
	expect := int64(1)
	actual := Panasonic2020B(1000000000, 1)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestPanasonic2020B5(t *testing.T) {
	expect := int64(1)
	actual := Panasonic2020B(1, 500000000)
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
