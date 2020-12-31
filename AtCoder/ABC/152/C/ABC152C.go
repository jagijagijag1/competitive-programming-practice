package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC152C() {
	sc.Split(bufio.ScanWords)
	N, P := nextInt(), []int{}
	for i := 0; i < N; i++ {
		P = append(P, nextInt())
	}

	fmt.Printf("%d\n", ABC152C(P))
}

// ABC152C ...
func ABC152C(P []int) (res int) {
	min := 1000000
	for _, p := range P {
		if p < min {
			res++
			min = p
		}
	}

	return
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
