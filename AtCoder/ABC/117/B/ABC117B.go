package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//func main() {
func mainABC117B() {
	sc.Split(bufio.ScanWords)
	N, L := nextInt(), []int{}
	for i := 0; i < N; i++ {
		L = append(L, nextInt())
	}

	fmt.Printf("%s\n", ABC117B(L))
}

// ABC117B ...
func ABC117B(L []int) string {
	sum, max := 0, -1
	for _, l := range L {
		if max < l {
			max = l
		}
		sum += l
	}

	if max < sum-max {
		return "Yes"
	}

	return "No"
}

// ABC117BNaive ...
func ABC117BNaive(L []int) string {
	maxIndex, maxVal := 0, -1
	for i, l := range L {
		if maxVal < l {
			maxIndex, maxVal = i, l
		}
	}
	L[maxIndex] = 0

	sum := 0
	for _, l := range L {
		sum += l
	}

	if maxVal < sum {
		return "Yes"
	}

	return "No"
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
