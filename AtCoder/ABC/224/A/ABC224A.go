package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	if strings.HasSuffix(nextLine(), "er") {
		fmt.Println("er")
	} else {
		fmt.Println("ist")
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