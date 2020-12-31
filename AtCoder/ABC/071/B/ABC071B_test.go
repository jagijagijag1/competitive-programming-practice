package main

import (
	"testing"
)

func TestABC071B1(t *testing.T) {
	expect := "b"
	actual := ABC071B("atcoderregularcontest")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC071B2(t *testing.T) {
	expect := "None"
	actual := ABC071B("abcdefghijklmnopqrstuvwxyz")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC071B3(t *testing.T) {
	expect := "d"
	actual := ABC071B("fajsonlslfepbjtsaayxbymeskptcumtwrmkkinjxnnucagfrg")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
