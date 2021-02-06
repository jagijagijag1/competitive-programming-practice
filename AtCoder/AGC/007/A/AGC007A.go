package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	H, W := nextInt(), nextInt()
	m, cnt := make([]string, H), 0
	for i := 0; i < H; i++ {
		m[i] = nextLine()
		for j := range m[i] {
			if m[i][j] == '#' {
				cnt++
			}
		}
	}

	if cnt == H+W-1 {
		fmt.Println("Possible")
	} else {
		fmt.Println("Impossible")
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
