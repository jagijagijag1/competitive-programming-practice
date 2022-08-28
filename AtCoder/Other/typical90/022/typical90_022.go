package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	a, b, c := nextInt(), nextInt(), nextInt()
	g := gcds(a, b, c)

	if g == 1 {
		fmt.Println(a - 1 + b - 1 + c - 1)
	} else {
		fmt.Println(a/g - 1 + b/g - 1 + c/g - 1)
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

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcds(a, b int, integers ...int) int {
	res := gcd(a, b)
	for i := 0; i < len(integers); i++ {
		res = gcd(res, integers[i])
	}
	return res
}

// var sc = bufio.NewScanner((os.Stdin))
