package main

import (
	"fmt"
)

func main() {
	var n, l, r int
	fmt.Scan(&n, &l, &r)
	printWithSeparator(1, l, r, "")
	for i := 2; i <= n; i++ {
		printWithSeparator(i, l, r, " ")
	}
	fmt.Println()
}

func printWithSeparator(i, l, r int, s string) {
	if i < l || i > r {
		fmt.Print(s, i)
	} else {
		fmt.Print(s, r-i+l)
	}
}
