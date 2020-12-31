package main

import (
	"testing"
)

func TestABC108B1(t *testing.T) {
	ex3, ey3, ex4, ey4 := -1, 1, -1, 0
	ax3, ay3, ax4, ay4 := ABC108B(0, 0, 0, 1)
	if ex3 != ax3 || ey3 != ay3 || ex4 != ax4 || ey4 != ay4 {
		t.Errorf(`expect="%d %d %d %d" actual="%d %d %d %d"`, ex3, ey3, ex4, ey4, ax3, ay3, ax4, ay4)
	}
}

func TestABC108B2(t *testing.T) {
	ex3, ey3, ex4, ey4 := 3, 10, -1, 7
	ax3, ay3, ax4, ay4 := ABC108B(2, 3, 6, 6)
	if ex3 != ax3 || ey3 != ay3 || ex4 != ax4 || ey4 != ay4 {
		t.Errorf(`expect="%d %d %d %d" actual="%d %d %d %d"`, ex3, ey3, ex4, ey4, ax3, ay3, ax4, ay4)
	}
}

func TestABC108B3(t *testing.T) {
	ex3, ey3, ex4, ey4 := -126, -64, -36, -131
	ax3, ay3, ax4, ay4 := ABC108B(31, -41, -59, 26)
	if ex3 != ax3 || ey3 != ay3 || ex4 != ax4 || ey4 != ay4 {
		t.Errorf(`expect="%d %d %d %d" actual="%d %d %d %d"`, ex3, ey3, ex4, ey4, ax3, ay3, ax4, ay4)
	}
}

func TestABC108B4(t *testing.T) {
	ex3, ey3, ex4, ey4 := 0, 6, -3, 3
	ax3, ay3, ax4, ay4 := ABC108B(0, 0, 3, 3)
	if ex3 != ax3 || ey3 != ay3 || ex4 != ax4 || ey4 != ay4 {
		t.Errorf(`expect="%d %d %d %d" actual="%d %d %d %d"`, ex3, ey3, ex4, ey4, ax3, ay3, ax4, ay4)
	}
}

func TestABC108B5(t *testing.T) {
	ex3, ey3, ex4, ey4 := 6, 0, 3, 3
	ax3, ay3, ax4, ay4 := ABC108B(0, 0, 3, -3)
	if ex3 != ax3 || ey3 != ay3 || ex4 != ax4 || ey4 != ay4 {
		t.Errorf(`expect="%d %d %d %d" actual="%d %d %d %d"`, ex3, ey3, ex4, ey4, ax3, ay3, ax4, ay4)
	}
}

func TestABC108B6(t *testing.T) {
	ex3, ey3, ex4, ey4 := -6, 0, -3, -3
	ax3, ay3, ax4, ay4 := ABC108B(0, 0, -3, 3)
	if ex3 != ax3 || ey3 != ay3 || ex4 != ax4 || ey4 != ay4 {
		t.Errorf(`expect="%d %d %d %d" actual="%d %d %d %d"`, ex3, ey3, ex4, ey4, ax3, ay3, ax4, ay4)
	}
}

func TestABC108B7(t *testing.T) {
	ex3, ey3, ex4, ey4 := 0, -6, 3, -3
	ax3, ay3, ax4, ay4 := ABC108B(0, 0, -3, -3)
	if ex3 != ax3 || ey3 != ay3 || ex4 != ax4 || ey4 != ay4 {
		t.Errorf(`expect="%d %d %d %d" actual="%d %d %d %d"`, ex3, ey3, ex4, ey4, ax3, ay3, ax4, ay4)
	}
}
