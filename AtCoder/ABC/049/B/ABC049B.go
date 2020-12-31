package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	H, _, C := nextInt(), nextInt(), []string{}
	for i := 0; i < H; i++ {
		C = append(C, nextLine())
	}

	res := make([]string, len(C)*2)
	for i := range C {
		res[2*i] = C[i]
		res[2*i+1] = C[i]
	}

	for _, r := range res {
		fmt.Println(r)
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
