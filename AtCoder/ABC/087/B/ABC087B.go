package main

import (
	"bufio"
	"fmt"
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
func mainABC087B() {
	sc.Split(bufio.ScanWords)
	a, b, c, x := nextInt(), nextInt(), nextInt(), nextInt()

	fmt.Printf("%d\n", ABC087B(a, b, c, x))
}

// ABC087B ...
func ABC087B(a, b, c, x int) int {
	res := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k == x {
					res++
				}
			}
		}
	}

	return res
}
