package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, W := nextInt(), []string{}
	for i := 0; i < N; i++ {
		W = append(W, nextLine())
	}

	used, lw := map[string]bool{}, ""
	for _, w := range W {
		if lw == "" {
			lw, used[w] = w, true
			continue
		}
		if lw[len(lw)-1] != w[0] || used[w] {
			fmt.Println("No")
			return
		}
		lw, used[w] = w, true
	}

	fmt.Println("Yes")
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
