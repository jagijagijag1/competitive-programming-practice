package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	S, Q := nextLine(), nextInt()

	rev, head, tail := false, []byte{}, []byte{}
	for i := 0; i < Q; i++ {
		t := nextInt()
		if t == 1 {
			rev = !rev
		} else {
			f, c := nextInt(), nextLine()[0]
			if f == 1 {
				if rev {
					tail = append(tail, c)
				} else {
					head = append(head, c)
				}
			} else {
				if rev {
					head = append(head, c)
				} else {
					tail = append(tail, c)
				}
			}
		}
	}

	headstr := reverse(string(head))
	S = headstr + S + string(tail)
	if rev {
		S = reverse((S))
	}
	fmt.Println(S)
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// var sc = bufio.NewScanner((os.Stdin))

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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)
