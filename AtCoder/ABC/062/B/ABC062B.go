package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	H, W, a := nextInt(), nextInt(), []string{}

	s := ""
	for i := 0; i < W+2; i++ {
		s += "#"
	}

	a = append(a, s)
	for i := 0; i < H; i++ {
		tmp := nextLine()
		a = append(a, "#"+tmp+"#")
	}
	a = append(a, s)

	for _, aa := range a {
		fmt.Println(aa)
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
