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
	n, res := nextInt(), 0

	for n%2 == 0 {
		n /= 2
		res++
	}
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			n /= i
			res++
		}
	}
	if n > 2 {
		res++
	}
	fmt.Println(int(math.Ceil(math.Log2(float64(res)))))
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
