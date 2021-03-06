package main

import (
	"bufio"
	"fmt"
	"math"
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
func mainABC079C() {
	sc.Split(bufio.ScanWords)
	ABCD := nextInt()
	var A, B, C, D int
	A, ABCD = ABCD/1000, ABCD%1000
	B, ABCD = ABCD/100, ABCD%100
	C, ABCD = ABCD/10, ABCD%10
	D = ABCD

	fmt.Printf("%s\n", ABC079C(A, B, C, D))
}

// ABC079C ...
func ABC079C(A, B, C, D int) string {
	results := make([]int, 16)
	resStr := make([]string, 16)
	results[0], resStr[0] = A, strconv.Itoa(A)
	rest := []int{A, B, C, D}

	for d := 0; d < len(rest)-1; d++ {
		dl := int(math.Pow(2, float64(d))) - 1
		for j := 0; j <= dl; j++ {
			i, n := dl+j, rest[d+1]
			results[2*i+1], resStr[2*i+1] = results[i]+n, resStr[i]+"+"+strconv.Itoa(n)
			results[2*i+2], resStr[2*i+2] = results[i]-n, resStr[i]+"-"+strconv.Itoa(n)
		}
	}

	for i := 7; i < len(results); i++ {
		if results[i] == 7 {
			return resStr[i] + "=7"
		}
	}

	return "something wrong"
}

// ABC079Cbitsearch ...
func ABC079Cbitsearch(A, B, C, D int) string {
	nums := []int{A, B, C, D}
	opNum := len(nums) - 1
	ops := make([]string, opNum)

	for i := 0; i < (1 << opNum); i++ {
		sum := nums[0]
		for j := 0; j < opNum; j++ {
			if (i>>j)&1 == 1 {
				ops[j] = "+"
				sum += nums[j+1]
			} else {
				ops[j] = "-"
				sum -= nums[j+1]
			}
		}

		if sum == 7 {
			break
		}
	}

	res := ""
	for i := 0; i < opNum; i++ {
		res += strconv.Itoa(nums[i]) + ops[i]
	}

	return res + strconv.Itoa(nums[len(nums)-1]) + "=7"
}
