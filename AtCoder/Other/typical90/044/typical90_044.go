package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, q := nextInt(), nextInt()
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	shift := 0
	for i := 0; i < q; i++ {
		t, x, y := nextInt(), (nextInt()-1-shift+n)%n, (nextInt()-1-shift+n)%n
		if t == 1 {
			a[x], a[y] = a[y], a[x]
		} else if t == 2 {
			shift = (shift + 1) % n
		} else {
			fmt.Println(a[x])
		}
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
