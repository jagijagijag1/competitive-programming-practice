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
	n, a := nextInt(), []tuple{}
	for i := 0; i < n; i++ {
		a = append(a, tuple{i + 1, nextInt()})
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].a < a[j].a {
			return true
		}
		return false
	})
	fmt.Println(a[len(a)-2].i)
}

type tuple struct {
	i, a int
}

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)

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

// var sc = bufio.NewScanner((os.Stdin))
