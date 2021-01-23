package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, C, K, T := nextInt(), nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		T = append(T, nextInt())
	}

	sort.Ints(T)
	res := 0
	for 0 < len(T) {
		h := T[0]
		T = T[1:]

		p := 1
		for p < C && 0 < len(T) {
			hh := T[0]
			if hh <= h+K {
				T = T[1:]
				p++
			} else {
				break
			}
		}
		res++
	}

	fmt.Println(res)
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
