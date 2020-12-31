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
	N, s, sum := nextInt(), []int{}, 0
	for i := 0; i < N; i++ {
		s = append(s, nextInt())
		sum += s[i]
	}

	if sum%10 != 0 {
		fmt.Println(sum)
		return
	}

	sort.Ints(s)
	for _, ss := range s {
		if ss%10 == 0 {
			continue
		}
		sum -= ss
		if sum%10 != 0 {
			fmt.Println(sum)
			return
		}
	}

	fmt.Println(0)
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
