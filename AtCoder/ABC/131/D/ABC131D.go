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
	N, t := nextInt(), []tuple{}
	for i := 0; i < N; i++ {
		t = append(t, tuple{nextInt(), nextInt()})
	}

	sort.Slice(t, func(i, j int) bool {
		if t[i].b < t[j].b {
			return true
		}
		return false
	})

	c := 0
	for _, tt := range t {
		c += tt.a
		if tt.b < c {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

type tuple struct {
	a, b int
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
