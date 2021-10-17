package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, x, y, res := nextInt(), []int{}, []int{}, 0

	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := float64(y[i]-y[j]) / float64(x[i]-x[j])
			if -1 <= d && d <= 1 {
				res++
			}
		}
	}
	fmt.Println(res)
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
