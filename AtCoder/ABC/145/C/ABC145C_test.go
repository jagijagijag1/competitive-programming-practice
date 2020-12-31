package main

import (
	"testing"
)

func TestABC145C1(t *testing.T) {
	expect := 2.2761423749
	actual := ABC145C(3, []Point{{0, 0}, {1, 0}, {0, 1}})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}

func TestABC145C2(t *testing.T) {
	expect := 91.9238815543
	actual := ABC145C(2, []Point{{-879, 981}, {-866, 890}})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}

func TestABC145C3(t *testing.T) {
	expect := 7641.9817824387
	actual := ABC145C(8, []Point{
		{-406, 10}, {512, 859}, {494, 362}, {-955, -475},
		{128, 553}, {-986, -885}, {763, 77}, {449, 310},
	})
	if actual != expect {
		t.Errorf(`expect="%f" actual="%f"`, expect, actual)
	}
}
