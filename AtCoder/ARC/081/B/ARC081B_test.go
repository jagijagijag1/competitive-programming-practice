package main

import (
	"testing"
)

func TestARC081B1(t *testing.T) {
	expect := 6
	actual := ARC081B(3, "aab", "ccb")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC081B2(t *testing.T) {
	expect := 3
	actual := ARC081B(1, "Z", "Z")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestARC081B3(t *testing.T) {
	expect := 958681902
	actual := ARC081B(52, "RvvttdWIyyPPQFFZZssffEEkkaSSDKqcibbeYrhAljCCGGJppHHn", "RLLwwdWIxxNNQUUXXVVMMooBBaggDKqcimmeYrhAljOOTTJuuzzn")
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
