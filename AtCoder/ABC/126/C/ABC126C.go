package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, K := nextInt(), nextInt()

	res := 0.0
	for i := 1; i <= N; i++ {
		score, cnt := i, 0
		for score < K {
			score *= 2
			cnt++
		}
		p := 1.0 / float64(N) * math.Pow(0.5, float64(cnt))
		res += p
	}

	fmt.Println(res)
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
