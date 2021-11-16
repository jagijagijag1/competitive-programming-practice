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
	n, k, f := nextInt(), nextInt(), []tuple{}
	for i := 0; i < n; i++ {
		f = append(f, tuple{nextInt(), nextInt()})
	}

	sort.Slice(f, func(i, j int) bool {
		if f[i].x < f[j].x {
			return true
		}
		return false
	})

	for _, ff := range f {
		if k < ff.x {
			break
		}
		k += ff.y
	}
	fmt.Println(k)
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

type tuple struct {
	x, y int
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
