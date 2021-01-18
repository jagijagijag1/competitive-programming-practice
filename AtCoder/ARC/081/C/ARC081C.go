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
	N, cnt, e := nextInt(), map[int]int{}, []int{}
	for i := 0; i < N; i++ {
		cnt[nextInt()]++
	}
	for k := range cnt {
		if cnt[k] >= 2 {
			e = append(e, k)
		}
	}

	if len(e) < 2 {
		fmt.Println(0)
	} else {
		sort.Ints(e)
		max := e[len(e)-1]
		if cnt[max] >= 4 {
			fmt.Println(max * max)
		} else {
			fmt.Println(max * e[len(e)-2])
		}
	}
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
