package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	A, B := nextInt(), nextInt()

	res := 0
	for i := A; i <= B; i++ {
		as := strconv.Itoa(i)
		if as[0] == as[4] && as[1] == as[3] {
			res++
		}
	}

	fmt.Printf("%d\n", res)
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
