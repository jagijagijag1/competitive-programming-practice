package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	x := nextInt()

	if x < 600 {
		fmt.Println(8)
	} else if x < 800 {
		fmt.Println(7)
	} else if x < 1000 {
		fmt.Println(6)
	} else if x < 1200 {
		fmt.Println(5)
	} else if x < 1400 {
		fmt.Println(4)
	} else if x < 1600 {
		fmt.Println(3)
	} else if x < 1800 {
		fmt.Println(2)
	} else {
		fmt.Println(1)
	}
}

var sc = bufio.NewScanner((os.Stdin))

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
