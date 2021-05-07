package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextLine()
	r := make([]rune, len(N))
	for i, n := range N {
		r[len(N)-i-1] = n
	}

	R := string(r)
	if R == N {
		fmt.Println("Yes")
		return
	}
	for i := 0; i < 10; i++ {
		N, R = "0"+N, R+"0"
		if R == N {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
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
