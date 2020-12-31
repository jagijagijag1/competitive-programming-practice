package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	l := nextLine()
	i, e := strconv.Atoi(l)
	if e != nil {
		panic(e)
	}
	return i
}

//func main() {
func mainABC070B() {
	sc.Split(bufio.ScanWords)
	A, B, C, D := nextInt(), nextInt(), nextInt(), nextInt()

	fmt.Printf("%d\n", ABC070B(A, B, C, D))
}

// ABC070B ...
func ABC070B(A, B, C, D int) int {
	var lower, upper int
	if A < C {
		lower = C
	} else {
		lower = A
	}

	if B < D {
		upper = B
	} else {
		upper = D
	}

	res := upper - lower
	if res < 0 {
		return 0
	}
	return res
}

// ABC070BNaiive ...
func ABC070BNaiive(A, B, C, D int) int {
	if B < C || D < A {
		return 0
	}

	if A <= C && D <= B {
		return D - C
	}
	if C <= A && B <= D {
		return B - A
	}

	if A < C {
		return B - C
	}

	return D - A
}
