package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, A, m := nextInt(), []int{}, map[int]int{}
	for i := 0; i < N; i++ {
		ai := nextInt()
		A = append(A, ai)
		m[ai]++
	}

	comb1, comb2, sum := map[int]int{}, map[int]int{}, 0
	for k, v := range m {
		if v == 1 {
			comb1[k] = 0
			comb2[k] = 0
		} else if v == 2 {
			comb1[k] = 1
			comb2[k] = 0
		} else {
			comb1[k] = v * (v - 1) / 2
			comb2[k] = (v - 1) * (v - 2) / 2
		}
		sum += comb1[k]
	}

	for _, a := range A {
		fmt.Println(sum - comb1[a] + comb2[a])
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
