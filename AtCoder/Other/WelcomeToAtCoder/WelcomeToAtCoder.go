package main

import (
	"fmt"
	"strconv"
)

func mainWelcomeToAtCoder() {
	var a, b, c int
	var s string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)
	// fmt.Printf("%d %s\n", a+b+c, s)
	fmt.Printf("%s\n", WelcomeToAtCoder(a, b, c, s))
}

// WelcomeToAtCoder ...
func WelcomeToAtCoder(a, b, c int, s string) string {
	return strconv.Itoa(a+b+c) + " " + s
}
