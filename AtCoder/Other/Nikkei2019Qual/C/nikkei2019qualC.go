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
	N, t := nextInt(), []triplet{}
	for i := 0; i < N; i++ {
		a, b := nextInt(), nextInt()
		t = append(t, triplet{a, b, a + b})
	}
	sort.Slice(t, func(i, j int) bool {
		if t[i].sum < t[j].sum {
			return true
		}
		return false
	})

	res, current := 0, 1
	for i := len(t) - 1; 0 <= i; i-- {
		if current == 1 {
			res += t[i].a
		} else {
			res -= t[i].b
		}
		current *= -1
	}
	fmt.Println(res)
}

type triplet struct {
	a, b, sum int
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
