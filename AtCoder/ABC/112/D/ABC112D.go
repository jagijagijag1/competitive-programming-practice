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
	N, M := nextInt(), nextInt()

	d := []int{1}
	for i := 1; i*i < M; i++ {
		if M%i == 0 {
			d = append(d, i)
			if i*i != M {
				d = append(d, M/i)
			}
		}
	}

	sort.Ints(d)
	idx := sort.Search(len(d), func(i int) bool {
		return d[i] > M/N
	})
	fmt.Println(d[max(0, idx-1)])
}

func divisor(n int) (res []int) {
	for i := 1; i*i < n; i++ {
		if n%i == 0 {
			res = append(res, i)
			if i*i != n {
				res = append(res, n/i)
			}
		}
	}
	sort.Ints(res)
	return
}

func max(a, b int) int {
	if a < b {
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
