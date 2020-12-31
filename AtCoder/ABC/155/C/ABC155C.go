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
	N := nextInt()
	resmap := map[string]int{}
	for i := 0; i < N; i++ {
		resmap[nextLine()]++
	}

	max, resslice := 0, []string{}
	for k, v := range resmap {
		if max < v {
			resslice = []string{}
			resslice = append(resslice, k)
			max = v
		} else if max == v {
			resslice = append(resslice, k)
		}
	}

	sort.Strings(resslice)
	for _, s := range resslice {
		fmt.Printf("%s\n", s)
	}
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
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
