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
	n, k := nextInt(), nextInt()
	a := make([][]int, k+1)
	for i := 0; i < n; i++ {
		a[i%k] = append(a[i%k], nextInt())
	}
	for i := 0; i < k; i++ {
		sort.Ints(a[i])
	}

	last := 0
	for j := 0; j < len(a[0]); j++ {
		for i := 0; i < k; i++ {
			if len(a[i]) <= j {
				break
			}
			if a[i][j] < last {
				fmt.Println("No")
				return
			}
			last = a[i][j]
		}
	}
	fmt.Println("Yes")
}

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

// var sc = bufio.NewScanner((os.Stdin))
