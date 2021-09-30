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
	a, b, h, m := float64(nextInt()), float64(nextInt()), float64(nextInt()), float64(nextInt())
	d := math.Abs((360.0/60.0)*m - math.Mod((360.0/(12.0*60.0))*(60*h+m), 360.0))
	fmt.Println(math.Sqrt(a*a + b*b - 2*a*b*math.Cos(d/180*math.Pi)))
}

func min(a, b float64) float64 {
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
