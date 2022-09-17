package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextLine(), nextInt()

	for i := 0; i < k; i++ {
		n10, d := 0, 1
		for j := 0; j < len(n); j++ {
			n10 += d * int(n[len(n)-1-j]-'0')
			d *= 8
		}

		n = ""
		for ; 9 <= n10; n10 /= 9 {
			n = strconv.Itoa(n10%9) + n
		}
		n = strconv.Itoa(n10) + n
		n = strings.ReplaceAll(n, "8", "5")
	}

	fmt.Println(n)
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
