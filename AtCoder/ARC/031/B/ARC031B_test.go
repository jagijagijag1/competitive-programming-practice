package main

import (
	"testing"
)

func TestARC031B1(t *testing.T) {
	expect := "YES"
	actual := ARC031B([]string{"xxxxxxxxxx", "xoooooooxx", "xxoooooxxx", "xxxoooxxxx", "xxxxoxxxxx", "xxxxxxxxxx", "xxxxoxxxxx", "xxxoooxxxx", "xxoooooxxx", "xxxxxxxxxx"})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC031B2(t *testing.T) {
	expect := "NO"
	actual := ARC031B([]string{"xxxxxxxxxx", "xoooooooxx", "xxoooooxxx", "xxxoooxxxx", "xxxxxxxxxx", "xxxxxxxxxx", "xxxxxxxxxx", "xxxoooxxxx", "xxoooooxxx", "xxxxxxxxxx"})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func TestARC031B3(t *testing.T) {
	expect := "YES"
	actual := ARC031B([]string{"xxxxoxxxxx", "xxxxoxxxxx", "xxxxoxxxxx", "xxxxoxxxxx", "ooooxooooo", "xxxxoxxxxx", "xxxxoxxxxx", "xxxxoxxxxx", "xxxxoxxxxx", "xxxxoxxxxx"})
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
