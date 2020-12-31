package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainARC144B() {
	sc.Split(bufio.ScanWords)
	N := nextInt()

	fmt.Printf("%s\n", ABC144B(N))
}

// ABC144B ...
func ABC144B(N int) string {
	for i := 1; i <= 9; i++ {
		if N%i == 0 {
			tmp := N / i
			if 1 <= tmp && tmp <= 9 {
				return "Yes"
			}
		}
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
