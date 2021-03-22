package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	a := make([]int, N+1)
	for i := 1; i <= N; i++ {
		a[i] = nextInt()
	}

	box, M := make([]int, N+1), 0
	for i := N; 0 < i; i-- {
		sum := 0
		for j := i; j <= N; j += i {
			sum += box[j]
		}

		if sum%2 != a[i] {
			box[i] = 1
			M++
		} else {
			box[i] = 0
		}
	}

	fmt.Println(M)
	if M == 0 {
		return
	}
	res := []int{}
	for i := 1; i <= N; i++ {
		if box[i] == 1 {
			res = append(res, i)
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
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
