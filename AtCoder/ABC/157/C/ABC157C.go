package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, M, s, c := nextInt(), nextInt(), []int{}, []int{}
	for i := 0; i < M; i++ {
		s = append(s, nextInt())
		c = append(c, nextInt())
	}

	res := make([]int, N)
	for i := range res {
		res[i] = -1
	}

	for i := range s {
		if res[s[i]-1] != -1 && res[s[i]-1] != c[i] {
			fmt.Println(-1)
			return
		}
		res[s[i]-1] = c[i]
	}

	if res[0] == 0 && N != 1 {
		fmt.Println(-1)
		return
	}

	for i := range res {
		if i == 0 && N != 1 {
			res[i] = max(res[i], 1)
		} else {
			res[i] = max(res[i], 0)
		}
		fmt.Print(res[i])
	}
	fmt.Println()
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
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
