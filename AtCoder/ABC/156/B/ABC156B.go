package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	fmt.Println(len(convertNtoBaseK(n, k)))
}

func convertNtoBaseK(n, k int) (res string) {
	tmp := n
	for digitNum := 1; ; {
		remainder := tmp % k
		if remainder >= 10 {
			res = string('A'+(remainder-10)) + res
		} else {
			res = string('0'+remainder) + res
		}
		tmp = tmp / k
		if tmp == 0 {
			break
		}
		digitNum *= 10
	}
	return
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
