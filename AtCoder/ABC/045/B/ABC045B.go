package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	cards := map[rune][]rune{}
	cards['a'] = []rune(nextLine())
	cards['b'] = []rune(nextLine())
	cards['c'] = []rune(nextLine())

	now := 'a'
	for {
		if len(cards[now]) == 0 {
			fmt.Println(string(now & 0x5f))
			return
		}
		now, cards[now] = cards[now][0], cards[now][1:]
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
