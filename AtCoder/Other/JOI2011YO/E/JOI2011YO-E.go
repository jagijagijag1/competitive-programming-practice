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
func mainJOI2011YOE() {
	sc.Split(bufio.ScanWords)
	H, W, N, c := nextInt(), nextInt(), nextInt(), []string{}
	for i := 0; i < H; i++ {
		c = append(c, nextLine())
	}

	fmt.Printf("%d\n", JOI2011YOE(H, W, N, c))
}

// JOI2011YOE ...
func JOI2011YOE(H, W, N int, c []string) (res int) {
	city, s := make([][]byte, H), Point{}
	for i := 0; i < H; i++ {
		city[i] = make([]byte, W)
		for j, char := range c[i] {
			city[i][j] = byte(char)
			if char == 'S' {
				s = Point{j, i}
			}
		}
	}

	cost, nextStart, vit := 0, s, 1
	for i := 0; i < N; i++ {
		cost, nextStart = bfsJOI2011YOE(city, nextStart, vit, H, W)
		res += cost
		vit++
	}

	return
}

func bfsJOI2011YOE(city [][]byte, s Point, vit, H, W int) (int, Point) {
	time := make([][]int, H)
	for i := 0; i < H; i++ {
		time[i] = make([]int, W)
		for j := 0; j < W; j++ {
			time[i][j] = 1000000
		}
	}

	var goal Point
	queue := []Point{s}
	time[s.y][s.x] = 0

	moves := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		for _, mv := range moves {
			next := Point{current.x + mv.x, current.y + mv.y}
			if next.x < 0 || next.y < 0 || W <= next.x || H <= next.y {
				continue
			}

			if city[next.y][next.x] == 'X' {
				continue
			} else if city[next.y][next.x] == '0'+byte(vit) {
				ct := time[current.y][current.x] + 1
				if ct < time[next.y][next.x] {
					time[next.y][next.x] = ct
					goal = Point{next.x, next.y}
				}
			} else {
				ct := time[current.y][current.x] + 1
				if ct < time[next.y][next.x] {
					time[next.y][next.x] = ct
					queue = append(queue, next)
				}
			}
		}
	}

	return time[goal.y][goal.x], goal
}
