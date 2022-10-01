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
	T, L, X, Y, Q := float64(nextInt()), float64(nextInt()), float64(nextInt()), float64(nextInt()), nextInt()

	for i := 0; i < Q; i++ {
		e := float64(nextInt())
		rad := math.Pi * (1.5 - 2*e/T)
		y, z := L/2*math.Cos(rad), L/2*math.Sin(rad)+L/2
		res := math.Atan2(z, math.Sqrt((y-Y)*(y-Y)+X*X))
		fmt.Printf("%.12f\n", res*180.0/math.Pi)
	}
}

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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)

// var sc = bufio.NewScanner((os.Stdin))
