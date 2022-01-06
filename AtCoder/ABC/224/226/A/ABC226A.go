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
	fmt.Println(math.Round(nextFloat64()))
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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextFloat64() float64 {
	l := nextLine()
	i, e := strconv.ParseFloat(l, 64)
	if e != nil {
		panic(e)
	}
	return i
}

// var sc = bufio.NewScanner((os.Stdin))
