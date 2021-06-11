package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	_, s, used := nextInt(), nextLine(), map[string]bool{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				r := regexp.MustCompile(`.*` + fmt.Sprint(i) + `.*` + fmt.Sprint(j) + `.*` + fmt.Sprint(k))
				p := fmt.Sprint(i) + fmt.Sprint(j) + fmt.Sprint(k)
				if r.MatchString(s) && !used[p] {
					used[p] = true
				}
			}
		}
	}
	fmt.Println(len(used))
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
