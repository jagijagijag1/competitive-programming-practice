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
	s, K, m := nextLine(), nextInt(), map[string]struct{}{}
	for i := 0; i < len(s); i++ {
		for k := 1; k <= K; k++ {
			m[s[i:min(i+k, len(s))]] = struct{}{}
		}
	}

	res := []string{}
	for k := range m {
		res = append(res, k)
	}
	sort.Strings(res)
	fmt.Println(res[K-1])
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
