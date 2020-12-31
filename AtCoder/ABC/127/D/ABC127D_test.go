package main

import (
	"testing"
)

func TestABC127D1(t *testing.T) {
	expect := 14
	actual := ABC127D([]int{5, 1, 4}, []ChangeCard{{2, 3}, {1, 5}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC127D2(t *testing.T) {
	expect := 338
	actual := ABC127D([]int{1, 8, 5, 7, 100, 4, 52, 33, 13, 5}, []ChangeCard{{3, 10}, {4, 30}, {1, 4}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC127D3(t *testing.T) {
	expect := 300
	actual := ABC127D([]int{100, 100, 100}, []ChangeCard{{3, 99}, {3, 99}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestABC127D4(t *testing.T) {
	expect := 10000000001
	actual := ABC127D([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []ChangeCard{{3, 1000000000}, {4, 1000000000}, {3, 1000000000}})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
