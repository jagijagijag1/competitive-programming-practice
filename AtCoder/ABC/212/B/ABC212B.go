package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	x := nextLine()

	if x[0] == x[1] && x[1] == x[2] && x[2] == x[3] {
		fmt.Println("Weak")
		return
	}

	if next(x[0]) == x[1] && next(x[1]) == x[2] && next(x[2]) == x[3] {
		fmt.Println("Weak")
		return
	}

	fmt.Println("Strong")
}

func next(x byte) byte {
	if x == '9' {
		return '0'
	} else {
		return x + 1
	}
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

// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

// var sc = bufio.NewScanner((os.Stdin))
