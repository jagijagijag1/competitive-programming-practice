package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, res := nextInt(), map[string]int{}

	res["AC"] = 0
	res["WA"] = 0
	res["TLE"] = 0
	res["RE"] = 0

	for i := 0; i < n; i++ {
		res[nextLine()]++
	}

	fmt.Println("AC x", res["AC"])
	fmt.Println("WA x", res["WA"])
	fmt.Println("TLE x", res["TLE"])
	fmt.Println("RE x", res["RE"])
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
