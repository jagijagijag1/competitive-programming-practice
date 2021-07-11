package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, narr := nextInt(), []int{}
	for i := 0; i < n; i++ {
		narr = append(narr, i+1)
	}
	perm := permutations(narr)

	p, q := []int{}, []int{}
	for i := 0; i < n; i++ {
		p = append(p, nextInt())
	}
	for i := 0; i < n; i++ {
		q = append(q, nextInt())
	}

	ps, qs, a, b := fmt.Sprint(p), fmt.Sprint(q), 0, 0
	for i := 0; i < len(perm); i++ {
		if perm[i] == ps {
			a = i + 1
		}
		if perm[i] == qs {
			b = i + 1
		}
	}
	fmt.Println(abs(a - b))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func permutations(arr []int) []string {
	var helper func([]int, int)
	res := []string{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, fmt.Sprint(tmp))
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	sort.Strings(res)
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

// const (
// 	initialBufSize = 10000
// 	maxBufSize     = 1000000
// )

// var (
// 	sc *bufio.Scanner = func() *bufio.Scanner {
// 		sc := bufio.NewScanner(os.Stdin)
// 		buf := make([]byte, initialBufSize)
// 		sc.Buffer(buf, maxBufSize)
// 		return sc
// 	}()
// )
