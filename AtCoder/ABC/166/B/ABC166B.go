package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	m := map[int]struct{}{}
	for i := 1; i <= n; i++ {
		m[i] = struct{}{}
	}
	for i := 0; i < k; i++ {
		d := nextInt()
		for j := 0; j < d; j++ {
			a := nextInt()
			if _, ok := m[a]; ok {
				delete(m, a)
			}
		}
	}
	fmt.Println(len(m))
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
