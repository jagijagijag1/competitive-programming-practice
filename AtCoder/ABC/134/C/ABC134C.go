package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainABC134C() {
	sc.Split(bufio.ScanWords)
	N, A := nextInt(), []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}

	// fmt.Printf("%d\n", ABC134C(A))
	for _, a := range ABC134C(A) {
		fmt.Printf("%d\n", a)
	}
}

// ABC134C ...
func ABC134C(A []int) (res []int) {
	AA := make([]int, len(A))
	copy(AA, A)
	sort.Ints(AA)
	for _, a := range A {
		if a == AA[len(A)-1] {
			res = append(res, AA[len(A)-2])
		} else {
			res = append(res, AA[len(A)-1])
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
