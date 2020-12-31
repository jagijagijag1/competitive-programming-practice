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
	N, K, m := nextInt(), nextInt(), map[int]tuple{}
	for i := 0; i < N; i++ {
		a := nextInt()
		if _, ok := m[a]; ok {
			m[a] = tuple{a, m[a].cnt + 1}
		} else {
			m[a] = tuple{a, 1}
		}
	}

	AA := []tuple{}
	for _, v := range m {
		AA = append(AA, v)
	}

	sort.Slice(AA, func(i, j int) bool {
		if AA[i].cnt < AA[j].cnt {
			return true
		}
		return false
	})

	res := 0
	for K < len(AA) {
		min := AA[0]
		AA = AA[1:]

		res += min.cnt
	}

	fmt.Println(res)
}

type tuple struct {
	num, cnt int
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
