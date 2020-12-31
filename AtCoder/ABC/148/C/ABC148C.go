package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC148C() {
	sc.Split(bufio.ScanWords)
	A, B := nextInt(), nextInt()

	fmt.Printf("%d\n", ABC148C(A, B))
}

// ABC148C ...
func ABC148C(A, B int) (res int) {
	return lcm(A, B)
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
