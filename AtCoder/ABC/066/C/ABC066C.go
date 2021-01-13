package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	res, c := make([]int, n), n/2
	res[c] = nextInt()
	for i := 1; i < n/2+1; i++ {
		if n%2 == 0 {
			res[c-i] = nextInt()
			if 2*i < n {
				res[c+i] = nextInt()
			}
		} else {
			res[c+i] = nextInt()
			res[c-i] = nextInt()
		}
	}

	// for i := range res {
	// 	fmt.Printf("%d", res[i])
	// 	if i != n-1 {
	// 		fmt.Print(" ")
	// 	}
	// }
	// fmt.Println()
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
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
