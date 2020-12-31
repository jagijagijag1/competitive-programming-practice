package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	X := float64(nextInt())

	if X == 1.0 {
		fmt.Println(1)
		return
	}

	max := 0.0
	for i := 2.0; i < X; i++ {
		for j := 2.0; ; j++ {
			res := math.Pow(i, j)
			if res > X {
				break
			} else if max < res {
				max = res
			}
		}
	}

	fmt.Println(max)
}

type triplet struct {
	i int
	S string
	P int
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
