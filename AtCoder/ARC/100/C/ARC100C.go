package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainARC100C() {
	sc.Split(bufio.ScanWords)
	N, A := nextInt(), []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}

	fmt.Printf("%d\n", ARC100C(A))
}

// ARC100C ...
func ARC100C(A []int) (res int) {
	zidx := []int{}
	for i, a := range A {
		zidx = append(zidx, a-(i+1))
	}
	sort.Sort(sort.IntSlice(zidx))
	mid := zidx[len(zidx)/2]
	fmt.Println(len(zidx)/2, ":", mid)

	for i, a := range A {
		res += absInt(a - (mid + i + 1))
	}

	// for k := -len(A); k <= len(A); k++ {
	// 	fmt.Printf("%d:: ", k)
	// 	ts := 0
	// 	for i, a := range A {
	// 		t := absInt(a - (k + i + 1))
	// 		fmt.Printf("%d ", t)
	// 		ts += t
	// 	}
	// 	fmt.Printf("= %d\n", ts)
	// }

	return
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
