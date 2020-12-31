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
	N, M, AB := nextInt(), nextInt(), []tuple{}
	for i := 0; i < N; i++ {
		AB = append(AB, tuple{nextInt(), nextInt()})
	}

	sort.Slice(AB, func(i, j int) bool {
		if AB[i].t1 < AB[j].t1 {
			return true
		}
		return false
	})

	num, cost := 0, 0
	for _, ab := range AB {
		if M <= num+ab.t2 {
			cost += ab.t1 * (M - num)
			break
		} else {
			cost += ab.t1 * ab.t2
			num += ab.t2
		}
	}

	fmt.Println(cost)
}

type tuple struct {
	t1, t2 int
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
