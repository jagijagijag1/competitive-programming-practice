package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainARC105B() {
	sc.Split(bufio.ScanWords)
	N := nextInt()

	fmt.Printf("%d\n", ABC105B(N))
}

// ABC105B ...
func ABC105B(N int) (res int) {
	for i := 1; i <= N; i += 2 {
		tmp := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				tmp++
			}
		}

		if tmp == 8 {
			res++
		}
	}

	return
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
