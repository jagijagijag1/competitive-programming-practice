package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, T, A := nextInt(), float64(nextInt()), float64(nextInt())
	minidx, mindiff := 0, 1000000000.0
	for i := 0; i < N; i++ {
		t := T - float64(nextInt())*0.006
		d := abs(t - A)
		if d < mindiff {
			minidx = i + 1
			mindiff = d
		}
	}

	fmt.Println(minidx)
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
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
