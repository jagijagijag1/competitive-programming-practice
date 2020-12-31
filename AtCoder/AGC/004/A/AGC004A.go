package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	A, B, C := nextInt(), nextInt(), nextInt()
	ABC := []int{A, B, C}
	sort.Ints(ABC)

	if A%2 == 0 || B%2 == 0 || C%2 == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(ABC[0] * ABC[1])
	}
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
