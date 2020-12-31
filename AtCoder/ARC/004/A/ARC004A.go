package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

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

// func main() {
func mainARC004A() {
	sc.Split(bufio.ScanWords)
	N, x, y := nextInt(), []int{}, []int{}
	for i := 0; i < N; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}

	fmt.Printf("%f\n", ARC004A(x, y))
}

// ARC004A ...
func ARC004A(x, y []int) (res float64) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			if i == j {
				continue
			}

			dx, dy := float64(x[i]-x[j]), float64(y[i]-y[j])
			tmp := math.Sqrt(dx*dx + dy*dy)
			if res < tmp {
				res = tmp
			}
		}
	}

	return
}
