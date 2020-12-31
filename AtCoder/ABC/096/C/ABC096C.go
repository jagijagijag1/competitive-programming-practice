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
func mainABC096C() {
	sc.Split(bufio.ScanWords)
	H, W, S := nextInt(), nextInt(), []string{}
	for i := 0; i < H; i++ {
		S = append(S, nextLine())
	}

	fmt.Printf("%s\n", ABC096C(H, W, S))
}

// ABC096C ...
func ABC096C(H, W int, S []string) string {
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, -1, 0, 1}

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
		for j := 1; j < W+1; j++ {
			if S[i][j] == '.' {
				continue
			}

			isAllWhite := true
			for d := 0; d < len(dx); d++ {
				if S[i+dx[d]][j+dy[d]] == '#' {
					isAllWhite = false
					break
				}
			}

			if isAllWhite {
				return "No"
			}

		}
	}

	return "Yes"
}
