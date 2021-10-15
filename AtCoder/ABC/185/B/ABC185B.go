package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m, t := nextInt(), nextInt(), nextInt()
	cn, ct := n, 0
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		cn -= a - ct
		if cn <= 0 {
			fmt.Println("No")
			return
		}

		cn = min(n, cn+(b-a))
		ct = b
		if t <= ct {
			break
		}
	}

	cn -= t - ct
	if cn <= 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
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
