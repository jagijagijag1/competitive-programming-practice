package main

import (
	"testing"
)

func TestABC075B1(t *testing.T) {
	ABC075B(3, 5, []string{".....", ".#.#.", "....."})
}

func TestABC075B2(t *testing.T) {
	ABC075B(3, 5, []string{"#####", "#####", "#####"})
}

func TestABC075B3(t *testing.T) {
	ABC075B(6, 6, []string{"#####.", "#.#.##", "####.#", ".#..#.", "#.##..", "#.#..."})
}
