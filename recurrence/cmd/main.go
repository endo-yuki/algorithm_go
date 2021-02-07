package main

import (
	"fmt"
)

func main() {
	a := 6
	fmt.Printf("1から%dまでの総和は%dです。", a, solve(a))
}

func solve(a int) int {
	if a == 0 {
		return 0
	} else {
		return a + solve(a-1)
	}
}
