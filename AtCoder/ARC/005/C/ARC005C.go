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

// Point ...
type Point struct {
	x int
	y int
}

// func main() {
func mainARC005C() {
	sc.Split(bufio.ScanWords)
	H, W, c := nextInt(), nextInt(), []string{}
	for i := 0; i < H; i++ {
		c = append(c, nextLine())
	}

	fmt.Printf("%s\n", ARC005C(H, W, c))
}

// ARC005C ...
func ARC005C(H, W int, c []string) string {
	queue, goal := []Point{}, Point{}
	board, cost := make([][]byte, H), make([][]int, H)
	for i := 0; i < H; i++ {
		board[i], cost[i] = make([]byte, W), make([]int, W)
		for j, char := range c[i] {
			board[i][j] = byte(char)
			cost[i][j] = 1000000
			if char == 's' {
				cost[i][j] = 0
				queue = append(queue, Point{j, i})
			}
			if char == 'g' {
				goal = Point{j, i}
			}
		}
	}

	moves := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		for _, mv := range moves {
			next := Point{current.x + mv.x, current.y + mv.y}
			if next.x < 0 || next.y < 0 || W <= next.x || H <= next.y {
				continue
			}

			ct := cost[current.y][current.x]
			if board[next.y][next.x] == '#' {
				ct++
			}
			if ct < cost[next.y][next.x] {
				cost[next.y][next.x] = ct
				queue = append(queue, next)
			}
		}
	}

	if cost[goal.y][goal.x] <= 2 {
		return "YES"
	}

	return "NO"
}
