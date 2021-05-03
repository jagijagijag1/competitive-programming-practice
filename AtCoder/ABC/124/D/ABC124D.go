package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, K, S := nextInt(), nextInt(), nextLine()
	nums, now, cnt := []int{}, 1, 0
	for i := 0; i < N; i++ {
		if S[i] == byte('0'+now) {
			cnt++
		} else {
			nums = append(nums, cnt)
			now, cnt = 1-now, 1
		}
	}
	if cnt != 0 {
		nums = append(nums, cnt)
	}
	if len(nums)%2 == 0 {
		nums = append(nums, 0)
	}

	res, add := 0, 2*K+1
	l, r, tmp := 0, 0, 0
	for i := 0; i < len(nums); i += 2 {
		nl, nr := i, min(i+add, len(nums))

		for l < nl {
			tmp -= nums[l]
			l++
		}
		for r < nr {
			tmp += nums[r]
			r++
		}

		res = max(tmp, res)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// var sc = bufio.NewScanner((os.Stdin))

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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)
