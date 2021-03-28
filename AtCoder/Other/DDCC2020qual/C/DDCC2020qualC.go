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
	H, W, K := nextInt(), nextInt(), nextInt()
	cake, res := make([][]rune, H), make([][]int, H)
	for i := 0; i < H; i++ {
		cake[i] = []rune(nextLine())
	}

	cnt := 0
	for h := 0; h < H; h++ {
		pos := []int{}
		for w := 0; w < W; w++ {
			if cake[h][w] == '#' {
				pos = append(pos, w)
			}
		}
		if len(pos) == 0 {
			continue
		}

		row, left := make([]int, W), 0
		for i, right := range pos {
			for ww := left; ww <= right; ww++ {
				row[ww] = cnt + i + 1
			}
			left = right + 1
		}
		for ww := left; ww < W; ww++ {
			row[ww] = cnt + len(pos)
		}
		cnt += len(pos)
		res[h] = row

		for hh := h - 1; 0 <= hh && len(res[hh]) == 0; hh-- {
			res[hh] = row
		}
		if cnt == K {
			for hh := h + 1; hh < H; hh++ {
				res[hh] = row
			}
			break
		}
	}

	for _, r := range res {
		fmt.Println(strings.Trim(fmt.Sprint(r), "[]"))
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
