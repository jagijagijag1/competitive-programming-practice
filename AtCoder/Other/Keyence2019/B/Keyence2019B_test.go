package main

import (
	"testing"
)

func TestKeyence2019B1(t *testing.T) {
	expect := "YES"
	actual := Keyence2019B("keyofscience")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestKeyence2019B2(t *testing.T) {
	expect := "NO"
	actual := Keyence2019B("mpyszsbznf")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestKeyence2019B3(t *testing.T) {
	expect := "NO"
	actual := Keyence2019B("ashlfyha")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestKeyence2019B4(t *testing.T) {
	expect := "YES"
	actual := Keyence2019B("keyence")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestKeyence2019B5(t *testing.T) {
	expect := "YES"
	actual := Keyence2019B("keyencea")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
