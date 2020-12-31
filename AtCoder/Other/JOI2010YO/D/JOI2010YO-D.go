package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
func mainJOI2010YOD() {
	sc.Split(bufio.ScanWords)
	n, k, cards := nextInt(), nextInt(), []string{}
	for i := 0; i < n; i++ {
		cards = append(cards, nextLine())
	}

	fmt.Printf("%d\n", JOI2010YOD(k, cards))
}

// JOI2010YOD ...
func JOI2010YOD(k int, cards []string) (res int) {
	results := map[string]struct{}{}

	for bit := 0; bit < (1 << uint(len(cards))); bit++ {
		if bits.OnesCount(uint(bit)) != k {
			continue
		}

		selected := []string{}
		for i := 0; i < len(cards); i++ {
			if bit>>uint(i)&1 == 1 {
				selected = append(selected, cards[i])
			}
		}

		prem := permutationsStr(selected)
		for _, p := range prem {
			tmp := ""
			for _, c := range p {
				tmp += c
			}
			results[tmp] = struct{}{}
		}

	}

	return len(results)
}

func permutationsStr(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
