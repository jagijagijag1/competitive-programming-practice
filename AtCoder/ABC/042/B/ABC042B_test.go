package main

import (
	"testing"
)

func TestABC042B1(t *testing.T) {
	expect := "axxcxxdxx"
	actual := ABC042B([]string{"dxx", "axx", "cxx"})
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
