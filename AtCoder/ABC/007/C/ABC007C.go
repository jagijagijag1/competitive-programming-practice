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
func mainABC007C() {
	sc.Split(bufio.ScanWords)
	R, C, sy, sx, gy, gx := nextInt(), nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	c := []string{}
	for i := 0; i < R; i++ {
		c = append(c, nextLine())
	}

	fmt.Printf("%d\n", ABC007C(R, C, sy, sx, gy, gx, c))
}

// ABC007C ...
func ABC007C(R, C, sy, sx, gy, gx int, c []string) int {
	maze, dist := make([][]byte, R), make([][]int, R)
	for i := 0; i < R; i++ {
		maze[i], dist[i] = make([]byte, C), make([]int, C)
		for j, char := range c[i] {
			maze[i][j] = byte(char)
			dist[i][j] = 10000
		}
	}

	s := Point{sx - 1, sy - 1}
	queue := []Point{s}
	dist[s.y][s.x] = 0

	moves := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		for _, mv := range moves {
			next := Point{current.x + mv.x, current.y + mv.y}
			if next.x < 0 || next.y < 0 || C <= next.x || R <= next.y {
				continue
			}
			if maze[next.y][next.x] == '#' {
				continue
			}

			cost := dist[current.y][current.x] + 1
			if cost < dist[next.y][next.x] {
				dist[next.y][next.x] = cost
				queue = append(queue, next)
			}
		}
	}

	return dist[gy-1][gx-1]
}
