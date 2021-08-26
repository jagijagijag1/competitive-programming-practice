package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextLine()

	k, res := len(n), 100000000
	for bit := 1; bit < (1 << uint(k)); bit++ {
		sum, cnt := 0, 0
		for i := 0; i < k; i++ {
			if bit>>uint(i)&1 == 1 {
				sum += int(n[i] - '0')
				cnt++
			}
		}
		if sum%3 == 0 {
			res = min(res, k-cnt)
		}
	}

	if res == 100000000 {
		res = -1
	}
	fmt.Println(res)
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

// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

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
