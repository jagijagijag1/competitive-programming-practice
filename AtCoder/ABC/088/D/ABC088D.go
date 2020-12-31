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
func mainABC088D() {
	sc.Split(bufio.ScanWords)
	H, W, s := nextInt(), nextInt(), []string{}
	for i := 0; i < H; i++ {
		s = append(s, nextLine())
	}

	fmt.Printf("%d\n", ABC088D(H, W, s))
}

// ABC088D ...
func ABC088D(H, W int, s []string) (res int) {
	maze, whitecnt := make([][]byte, H), 0
	for i := 0; i < H; i++ {
		maze[i] = make([]byte, W)
		for j, char := range s[i] {
			maze[i][j] = byte(char)
			if char == '.' {
				whitecnt++
			}
		}
	}

	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = 1000000
		}
	}

	start, goal := Point{0, 0}, Point{W - 1, H - 1}
	queue := []Point{start}
	dist[start.y][start.x] = 0

	moves := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.y == goal.y && current.x == goal.x {
			continue
		}

		for _, mv := range moves {
			next := Point{current.x + mv.x, current.y + mv.y}
			if next.x < 0 || next.y < 0 || W <= next.x || H <= next.y {
				continue
			}

			if maze[next.y][next.x] == '#' {
				continue
			} else {
				nextdist := dist[current.y][current.x] + 1
				if nextdist < dist[next.y][next.x] {
					dist[next.y][next.x] = nextdist
					queue = append(queue, next)
				}
			}
		}
	}

	if dist[goal.y][goal.x] >= 1000000 {
		return -1
	}

	return whitecnt - (dist[goal.y][goal.x] + 1)
}
