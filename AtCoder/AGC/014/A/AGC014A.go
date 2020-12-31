package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainAGC014A() {
	sc.Split(bufio.ScanWords)
	A, B, C := nextInt(), nextInt(), nextInt()

	fmt.Printf("%d\n", AGC014A(A, B, C))
}

// AGC014A ...
func AGC014A(A, B, C int) (res int) {
	la, lb, lc := A, B, C
	for {
		if A%2 == 1 || B%2 == 1 || C%2 == 1 {
			return res
		}

		la, lb, lc = A, B, C
		A, B, C = (B+C)/2, (A+C)/2, (A+B)/2

		if la >= A && lb >= B && lc >= C {
			return -1
		}

		res++
	}
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
