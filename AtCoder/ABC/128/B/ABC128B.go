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
	N, SP := nextInt(), []triplet{}
	for i := 0; i < N; i++ {
		SP = append(SP, triplet{i + 1, nextLine(), nextInt()})
	}

	sort.Slice(SP, func(i, j int) bool {
		if SP[i].S < SP[j].S {
			return true
		} else if SP[i].S == SP[j].S {
			if SP[i].P > SP[j].P {
				return true
			}
		}
		return false
	})

	for _, sp := range SP {
		fmt.Println(sp.i)
	}
}

type triplet struct {
	i int
	S string
	P int
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
