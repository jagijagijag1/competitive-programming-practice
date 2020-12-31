package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

//func main() {
func mainABC075B() {
	sc.Split(bufio.ScanWords)
	H, W, S := nextInt(), nextInt(), []string{}
	for i := 0; i < H; i++ {
		S = append(S, nextLine())
	}

	ABC075B(H, W, S)
}

// ABC075B ...
func ABC075B(H, W int, S []string) {
	op := []int{-1, 0, 1}

	var guard string
	for i := 0; i < W; i++ {
		guard += "."
	}
	S = append([]string{guard}, S...)
	S = append(S, guard)

	for i := 0; i < len(S); i++ {
		S[i] = "." + S[i] + "."
	}

	for i := 1; i < H+1; i++ {
		res := ""
		for j := 1; j < W+1; j++ {
			if S[i][j] == '#' {
				res += "#"
			} else {
				num := 0
				for _, dx := range op {
					for _, dy := range op {
						if S[i+dx][j+dy] == '#' {
							num++
						}
					}
				}
				res += strconv.Itoa(num)
			}
		}
		fmt.Println(res)
	}
}
