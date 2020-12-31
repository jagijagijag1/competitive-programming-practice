package main

import (
	"testing"
)

func TestABC058B1(t *testing.T) {
	expect := "xaybzc"
	actual := ABC058B("xyz", "abc")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestABC058B2(t *testing.T) {
	expect := "aattccooddeerrbreeggiunlnaerrccoonntteesstt"
	actual := ABC058B("atcoderbeginnercontest", "atcoderregularcontest")
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
