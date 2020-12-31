package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, M, pS := nextInt(), nextInt(), []tuple{}
	for i := 0; i < M; i++ {
		pS = append(pS, tuple{nextInt(), nextLine()})
	}

	ACs, WAs, nac, np := make([]bool, N+1), make([]int, N+1), 0, 0
	for i := range pS {
		if !ACs[pS[i].p] && pS[i].S == "AC" {
			ACs[pS[i].p] = true
			nac++
		} else if !ACs[pS[i].p] && pS[i].S == "WA" {
			WAs[pS[i].p]++
		}
	}

	for i := range ACs {
		if ACs[i] {
			np += WAs[i]
		}
	}

	fmt.Println(nac, np)
}

type tuple struct {
	p int
	S string
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
