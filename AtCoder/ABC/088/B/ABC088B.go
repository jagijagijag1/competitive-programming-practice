package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
func mainABC088B() {
	sc.Split(bufio.ScanWords)
	N, a := nextInt(), []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	fmt.Printf("%d\n", ABC088B(a))
}

// ABC088B ...
func ABC088B(a []int) int {
	res := 0

	sort.Ints(a)
	for i := len(a) - 1; i >= 0; i -= 2 {
		res += a[i]
		if i != 0 {
			res -= a[i-1]
		}
	}

	return res
}
