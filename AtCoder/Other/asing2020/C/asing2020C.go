package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	res := make([]int, n+1)
	for x := 1; x*x <= n; x++ {
		for y := 1; y*y <= n; y++ {
			for z := 1; z*z <= n; z++ {
				nn := x*x + y*y + z*z + x*y + y*z + z*x
				if nn <= n {
					res[nn]++
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Println(res[i])
	}
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
