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

// func main() {
func mainARC031B() {
	sc.Split(bufio.ScanWords)
	A := []string{}
	for i := 0; i < 10; i++ {
		A = append(A, nextLine())
	}

	fmt.Printf("%s\n", ARC031B(A))
}

// ARC031B ...
func ARC031B(A []string) string {
	m, checked := make([][]byte, 10), make([][]bool, 10)
	initMaps(A, m, checked)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			m[i][j] = 'o'

			fillcnt := 0
			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {
					if m[k][l] == 'o' {
						fillcnt++
						if fillcnt > 1 {
							break
						} else {
							m, checked = dfsARC031B(k, l, m, checked)
						}
					}
				}
				if fillcnt > 1 {
					break
				}
			}

			if fillcnt == 1 {
				return "YES"
			}
			initMaps(A, m, checked)
		}
	}

	return "NO"
}

func dfsARC031B(x, y int, m [][]byte, checked [][]bool) ([][]byte, [][]bool) {
	if x < 0 || y < 0 || 10 <= x || 10 <= y {
		return m, checked
	}
	if checked[x][y] {
		return m, checked
	}
	if m[x][y] == 'x' {
		return m, checked
	}
	if m[x][y] == 'o' {
		m[x][y] = 'x'
	}

	checked[x][y] = true
	dx, dy := []int{1, 0, -1, 0}, []int{0, 1, 0, -1}

	for i := 0; i < len(dx); i++ {
		nx, ny := x+dx[i], y+dy[i]
		m, checked = dfsARC031B(nx, ny, m, checked)
	}

	return m, checked
}

func initMaps(A []string, m [][]byte, checked [][]bool) {
	for i := 0; i < 10; i++ {
		m[i], checked[i] = make([]byte, 10), make([]bool, 10)
		for j, char := range A[i] {
			m[i][j] = byte(char)
			checked[i][j] = false
		}
	}
}
