package main

import (
	"fmt"
)

// var sc = bufio.NewScanner((os.Stdin))
//
// func nextLine() string {
// 	sc.Scan()
// 	return sc.Text()
// }
//
// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

// func main() {
func mainATC001A() {
	// sc.Split(bufio.ScanWords)
	// H, W, c := nextInt(), nextInt(), []string{}
	// for i := 0; i < H; i++ {
	// 	c = append(c, nextLine())
	// }

	// fmt.Printf("%s\n", ATC001A(H, W, c))

	var h, w int
	fmt.Scan(&h, &w)

	var sx, sy int
	graph, visited := make([][]byte, h), make([][]bool, h)

	for i := 0; i < h; i++ {
		graph[i] = make([]byte, w)
		visited[i] = make([]bool, w)
		var s []byte
		fmt.Scan(&s)
		for j := 0; j < w; j++ {
			graph[i][j] = s[j]
			if s[j] == 's' {
				sx = i
				sy = j
			}
		}
	}

	if dfsATC001A2(sx, sy, h, w, graph, visited) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfsATC001A2(x, y, h, w int, maze [][]byte, visited [][]bool) bool {
	if x < 0 || x >= h || y < 0 || y >= w {
		return false
	}

	if visited[x][y] {
		return false
	}
	visited[x][y] = true

	if maze[x][y] == 'g' {
		return true
	}
	if maze[x][y] == '#' {
		return false
	}

	for _, a := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nx := x + a[0]
		ny := y + a[1]
		if dfsATC001A2(nx, ny, h, w, maze, visited) {
			return true
		}
	}
	return false
}

// ATC001A ...
func ATC001A(H, W int, c []string) string {
	m, s, checked := make([][]string, H), Point{}, make([][]bool, H)
	for i := 0; i < H; i++ {
		m[i], checked[i] = make([]string, W), make([]bool, W)
		for j, char := range c[i] {
			m[i][j] = string(char)
			checked[i][j] = false
			if m[i][j] == "s" {
				s = Point{i, j}
			}
		}
	}

	if dfsATC001A(s, H, W, m, checked) {
		return "Yes"
	}
	return "No"
}

// Point ...
type Point struct {
	x int
	y int
}

func dfsATC001A(current Point, H, W int, m [][]string, checked [][]bool) bool {
	moves := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	if current.x < 0 || current.y < 0 || W <= current.x || H <= current.y {
		return false
	} else if checked[current.y][current.x] {
		return false
	} else if m[current.y][current.x] == "#" {
		return false
	} else if m[current.y][current.x] == "g" {
		return true
	}

	checked[current.y][current.x] = true

	for _, mv := range moves {
		next := Point{current.x + mv.x, current.y + mv.y}
		if dfsATC001A(next, H, W, m, checked) {
			return true
		}
	}

	return false
}
