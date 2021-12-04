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
	n, k, a := nextInt(), nextInt(), []triple{}
	for i := 0; i < n; i++ {
		a = append(a, triple{i + 1, nextInt(), 0})
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].ai < a[j].ai {
			return true
		}
		return false
	})

	d := 0
	if n <= k {
		d += k / n
		k %= n
	}

	for i := 0; i < n; i++ {
		if i < k {
			a[i].n++
		}
		a[i].n += d
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].i < a[j].i {
			return true
		}
		return false
	})

	out := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d\n", a[i].n)
	}
	defer out.Flush()
}

type triple struct {
	i, ai, n int
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
