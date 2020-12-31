package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainARC068B() {
	sc.Split(bufio.ScanWords)
	N := nextInt()

	fmt.Printf("%d\n", ABC068B(N))
}

// ABC068B ...
func ABC068B(N int) int {
	maxNum, maxDiv := -1, -1
	for i := 1; i <= N; i++ {
		tmpNum, tmpDiv := i, 0
		for {
			if tmpNum%2 == 0 {
				tmpNum /= 2
				tmpDiv++
			} else {
				break
			}
		}

		if maxDiv < tmpDiv {
			maxNum = i
			maxDiv = tmpDiv
		}
	}

	return maxNum
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
