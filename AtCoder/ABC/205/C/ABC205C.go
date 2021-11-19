package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	a, b, c := nextInt(), nextInt(), nextInt()

	if c%2 == 0 {
		c = 2
	} else {
		c = 1
	}

	ac, bc := pow(a, c), pow(b, c)
	if ac == bc {
		fmt.Println("=")
	} else if ac < bc {
		fmt.Println("<")
	} else {
		fmt.Println(">")
	}

}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
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
