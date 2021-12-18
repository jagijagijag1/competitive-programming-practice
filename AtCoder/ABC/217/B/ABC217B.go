package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	m := map[string]struct{}{"ABC": {}, "ARC": {}, "AGC": {}, "AHC": {}}
	for i := 0; i < 3; i++ {
		delete(m, nextLine())
	}
	fmt.Println(strings.Trim(fmt.Sprint(m), "map[:{}]"))
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
