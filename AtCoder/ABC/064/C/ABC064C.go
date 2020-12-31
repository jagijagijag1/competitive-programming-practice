package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC064C() {
	sc.Split(bufio.ScanWords)
	N, a := nextInt(), []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	r1, r2 := ABC064C(a)
	fmt.Printf("%d %d\n", r1, r2)
}

// ABC064C ...
func ABC064C(a []int) (int, int) {
	rate := [9]int{}
	for _, at := range a {
		if at >= 3200 {
			rate[8]++
		} else {
			rate[at/400] = 1
		}
	}

	min := 0
	for i := 0; i < 8; i++ {
		min += rate[i]
	}

	max := min + rate[8]

	if min == 0 {
		min++
	}

	return min, max
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
