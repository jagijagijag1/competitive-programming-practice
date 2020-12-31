package main

import (
	"testing"
)

func TestABC165C1(t *testing.T) {
	expect := 110
	actual := ABC165C(3, 4, []abcd{{1, 3, 3, 100}, {1, 2, 2, 10}, {2, 3, 2, 10}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC165C2(t *testing.T) {
	expect := 357500
	actual := ABC165C(4, 6, []abcd{{2, 4, 1, 86568}, {1, 4, 0, 90629}, {2, 3, 0, 90310}, {3, 4, 1, 29211}, {3, 4, 3, 78537}, {3, 4, 2, 8580}, {1, 2, 1, 96263}, {1, 4, 2, 2156}, {1, 2, 0, 94325}, {1, 4, 3, 94328}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC165C3(t *testing.T) {
	expect := 1
	actual := ABC165C(10, 10, []abcd{{1, 10, 9, 1}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
