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
	n := nextInt()
	om, em := map[int]int{}, map[int]int{}
	for i := 1; i <= n; i++ {
		v := nextInt()
		if i%2 != 0 {
			om[v]++
		} else {
			em[v]++
		}
	}

	os := []tuple{}
	for k, v := range om {
		os = append(os, tuple{k, v})
	}
	sort.Slice(os, func(i, j int) bool {
		return os[i].t2 > os[j].t2
	})

	es := []tuple{}
	for k, v := range em {
		es = append(es, tuple{k, v})
	}
	sort.Slice(es, func(i, j int) bool {
		return es[i].t2 > es[j].t2
	})

	if len(os) == 1 && len(es) == 1 && os[0].t1 == es[0].t1 {
		fmt.Println(min(os[0].t2, es[0].t2))
	} else if os[0].t1 == es[0].t1 {
		res1 := n - os[0].t2 - es[1].t2
		res2 := n - os[1].t2 - es[0].t2
		fmt.Println(min(res1, res2))
	} else {
		fmt.Println(n - os[0].t2 - es[0].t2)
	}
}

type tuple struct {
	t1, t2 int
}

func min(a, b int) int {
	if a > b {
		return b
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
