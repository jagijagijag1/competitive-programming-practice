package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	A, B, C, D := nextInt()-1, nextInt(), nextInt(), nextInt()
	CD := lcm(C, D)

	AC, AD, ACD := A/C, A/D, A/CD
	BC, BD, BCD := B/C, B/D, B/CD
	fmt.Println((B - BC - BD + BCD) - (A - AC - AD + ACD))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b) // aとbの2数ならここだけ

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

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
