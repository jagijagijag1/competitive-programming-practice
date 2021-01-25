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
	N, K, S := nextInt(), nextInt(), nextInt()

	A := make([]int, N)
	var rest int
	if S == 1000000000 {

		rest = 1
	} else {
		rest = S + 1
	}

	for i := 0; i < K; i++ {
		A[i] = S
	}
	for i := K; i < N; i++ {
		A[i] = rest
	}

	fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
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

// var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

// func readLine() string {
// 	buf := make([]byte, 0, 1000000)
// 	for {
// 		l, p, e := rdr.ReadLine()
// 		if e != nil {
// 			panic(e)
// 		}
// 		buf = append(buf, l...)
// 		if !p {
// 			break
// 		}
// 	}
// 	return string(buf)
// }
