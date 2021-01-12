package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	X := nextInt()

	sum := 0
	for i := 1; i <= X; i++ {
		l := X - i
		if l <= sum {
			fmt.Println(i)
			return
		}
		sum += i
	}

	fmt.Println(X)
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
