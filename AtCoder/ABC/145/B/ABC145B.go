package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	_, s := nextLine(), nextLine()
	if s[:len(s)/2] == s[len(s)/2:] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
