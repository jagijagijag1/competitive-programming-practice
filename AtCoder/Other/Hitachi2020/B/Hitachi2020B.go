package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// func main() {
func mainHitachi2020B() {
	sc.Split(bufio.ScanWords)
	A, B, M := nextInt(), nextInt(), nextInt()
	a, b, coupons := []int{}, []int{}, []coupon{}

	for i := 0; i < A; i++ {
		a = append(a, nextInt())
	}
	for i := 0; i < B; i++ {
		b = append(b, nextInt())
	}
	for i := 0; i < M; i++ {
		coupons = append(coupons, coupon{nextInt(), nextInt(), nextInt()})
	}

	fmt.Printf("%d\n", Hitachi2020B(a, b, coupons))
}

// Hitachi2020B ...
func Hitachi2020B(a, b []int, coupons []coupon) int {
	amin := minIntInSlice(a)
	bmin := minIntInSlice(b)

	res := amin + bmin
	for _, c := range coupons {
		t := a[c.x-1] + b[c.y-1] - c.c
		if t < res {
			res = t
		}
	}

	return res
}

type coupon struct {
	x, y, c int
}

func minIntInSlice(v []int) (res int) {
	res = math.MaxInt32
	for i, e := range v {
		if i == 0 || res > e {
			res = e
		}
	}
	return
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
