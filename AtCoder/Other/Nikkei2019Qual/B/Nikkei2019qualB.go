package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainNikkei2019qualB() {
	sc.Split(bufio.ScanWords)
	N, D := nextInt(), []int{}
	for i := 0; i < N; i++ {
		D = append(D, nextInt())
	}

	fmt.Printf("%d\n", Nikkei2019qualB(D))
}

// Nikkei2019qualB ...
func Nikkei2019qualB(D []int) int {
	if D[0] != 0 {
		return 0
	}

	mod := 998244353
	sort.Sort(sort.IntSlice(D))

	maxd, ap := 0, map[int]int{}
	for _, a := range D {
		ap[a]++
		if maxd < a {
			maxd = a
		}
	}

	if ap[0] != 1 {
		return 0
	}

	res := 1
	for i := 1; i <= maxd; i++ {
		v, ok := ap[i]
		if !ok {
			return 0
		}

		for j := 0; j < v; j++ {
			res = (res * ap[i-1]) % mod
		}
	}

	return res
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
