package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, a, m := nextInt(), []int{}, 0
	for i := 0; i < n; i++ {
		aa := nextInt()
		a = append(a, aa)
		m = max(m, aa)
	}

	res, cnt := 0, 0
	for i := 2; i <= m; i++ {
		tmp := 0
		for j := 0; j < n; j++ {
			if a[j]%i == 0 {
				tmp++
			}
		}
		if cnt < tmp {
			res = i
			cnt = tmp
		}
	}
	fmt.Println(res)
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
