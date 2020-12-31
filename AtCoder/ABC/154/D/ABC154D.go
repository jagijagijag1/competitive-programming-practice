package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, K := nextInt(), nextInt()
	e, c := make([]float64, N), make([]float64, N)

	e[0] = float64(nextInt()+1) / 2.0
	c[0] = e[0]
	res := c[0]
	for i := 1; i < N; i++ {
		e[i] = float64(nextInt()+1) / 2.0
		c[i] = c[i-1] + e[i]
		tmp := 0.0
		if i < K {
			tmp = c[i]
		} else {
			tmp = c[i-1] + e[i] - c[i-K]
		}
		if res < tmp {
			res = tmp
		}
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
