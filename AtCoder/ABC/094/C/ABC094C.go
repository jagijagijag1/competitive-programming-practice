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
	N, X := nextInt(), []triplet{}
	for i := 0; i < N; i++ {
		X = append(X, triplet{nextInt(), i, -1})
	}

	sort.Slice(X, func(i, j int) bool {
		if X[i].num < X[j].num {
			return true
		}
		return false
	})
	for i := range X {
		X[i].sortedidx = i
	}
	m1, m2 := X[N/2-1].num, X[N/2].num

	sort.Slice(X, func(i, j int) bool {
		if X[i].idx < X[j].idx {
			return true
		}
		return false
	})
	for _, x := range X {
		if x.sortedidx < N/2 {
			fmt.Println(m2)
		} else {
			fmt.Println(m1)
		}
	}
}

type triplet struct {
	num, idx, sortedidx int
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
