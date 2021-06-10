package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	sc.Split(bufio.ScanWords)
	s := nextLine()
	r := regexp.MustCompile(`^[A]{0,1}KIH[A]{0,1}B[A]{0,1}R[A]{0,1}$`)
	if r.MatchString(s) {
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
