package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	s, t := []string{nextLine(), nextLine(), nextLine()}, nextLine()
	out := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(t); i++ {
		fmt.Fprintf(out, "%s", s[int(t[i]-'0')-1])
	}
	fmt.Fprintf(out, "\n")
	defer out.Flush()
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
