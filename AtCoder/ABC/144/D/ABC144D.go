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
	a, b, x := float64(nextInt()), float64(nextInt()), float64(nextInt())

	aa := (2.0 * x) / (a * b)
	if aa <= a {
		r := math.Atan(b / aa)
		fmt.Println(r * 180.0 / math.Pi)
	} else {
		aa = 2 * (b - x/(a*a))
		r := math.Atan(aa / a)
		fmt.Println(r * 180.0 / math.Pi)
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
