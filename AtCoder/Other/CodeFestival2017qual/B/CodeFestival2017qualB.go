package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, candidates := nextInt(), map[int]int{}
	for i := 0; i < N; i++ {
		candidates[nextInt()]++
	}

	M := nextInt()
	for i := 0; i < M; i++ {
		mi := nextInt()
		r := candidates[mi]
		if r == 0 {
			fmt.Println("NO")
			return
		}
		candidates[mi]--
	}

	fmt.Println("YES")
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
