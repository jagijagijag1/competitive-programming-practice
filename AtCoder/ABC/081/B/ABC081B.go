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

//func main() {
func mainABC081B() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := []int{}
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
	}

	fmt.Printf("%d\n", ABC081B(a))
}

// ABC081B ...
func ABC081B(a []int) int {
	res := math.MaxInt8

	for _, n := range a {
		tmp := 0
		for n%2 == 0 {
			n = n >> 1
			tmp++
		}

		if tmp < res {
			res = tmp
		}
	}

	return res
}

// ABC081Bnaiive ...
func ABC081Bnaiive(a []int) int {
	res := 0

	f := true
	for f {
		for i, n := range a {
			if n%2 == 1 {
				f = false
				break
			}
			a[i] = n / 2
		}

		if f {
			res++
		}
	}

	return res
}
