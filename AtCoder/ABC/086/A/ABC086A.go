package main

import (
	"fmt"
)

//func main() {
func mainABC086A() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%s\n", ABC086A(a, b))
}

// ABC086A ...
func ABC086A(a, b int) string {
	if a%2 == 0 || b%2 == 0 {
		return "Even"
	}
	return "Odd"
}
