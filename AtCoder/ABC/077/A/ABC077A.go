package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	c1, c2 := nextLine(), nextLine()
	if c1[0] == c2[2] && c1[1] == c2[1] && c1[2] == c2[0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

var sc = bufio.NewScanner((os.Stdin))

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
