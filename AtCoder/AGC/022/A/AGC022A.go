package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	sc.Split(bufio.ScanWords)
	S, used := []rune(nextLine()), map[rune]bool{}
	for _, s := range S {
		used[s] = true
	}

	if len(S) == 26 {
		sidx, suffix := 0, []string{}
		for i := len(S) - 2; 0 <= i; i-- {
			suffix = append(suffix, string(S[i+1]))
			if S[i] < S[i+1] {
				sidx = i + 1
				break
			}
		}
		if sidx == 0 {
			fmt.Println(-1)
		} else {
			sort.Strings(suffix)
			for _, s := range suffix {
				if string(S[sidx-1]) < s {
					fmt.Println(string(S[0:sidx-1]) + s)
					break
				}
			}
		}
	} else {
		for i := 'a'; i <= 'z'; i++ {
			if !used[i] {
				S = append(S, i)
				fmt.Println(string(S))
				return
			}
		}
	}
}

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

// const (
// 	initialBufSize = 10000
// 	maxBufSize     = 1000000
// )

// var (
// 	sc *bufio.Scanner = func() *bufio.Scanner {
// 		sc := bufio.NewScanner(os.Stdin)
// 		buf := make([]byte, initialBufSize)
// 		sc.Buffer(buf, maxBufSize)
// 		return sc
// 	}()
// )
