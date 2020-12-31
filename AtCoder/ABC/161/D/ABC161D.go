package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC161D() {
	sc.Split(bufio.ScanWords)
	K := nextInt()

	fmt.Printf("%d\n", ABC161D(K))
}

// ABC161D ...
func ABC161D(K int) uint64 {
	queue := []uint64{}

	for i := 1; i <= 9; i++ {
		if i == K {
			return uint64(i)
		}
		queue = append(queue, uint64(i))
	}

	cnt := 9
	for {
		q := queue[0]
		queue = queue[1:]

		if q%10 != 0 {
			next1 := q*10 + ((q % 10) - 1)
			cnt++
			if cnt == K {
				return next1
			}
			queue = append(queue, next1)
		}

		next2 := q*10 + (q % 10)
		cnt++
		if cnt == K {
			return next2
		}
		queue = append(queue, next2)

		if q%10 != 9 {
			next3 := q*10 + ((q % 10) + 1)
			cnt++
			if cnt == K {
				return next3
			}
			queue = append(queue, next3)
		}

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
