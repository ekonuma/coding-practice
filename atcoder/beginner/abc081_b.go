package main

import (
	"fmt"
)

func main() {
	var n int
	var a []int

	fmt.Scanf("%d", &n)
	a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	count := -1
	do := true
	for do {
		count++
		for i, an := range a {
			if an%2 == 0 {
				a[i] = an / 2
			} else {
				do = false
				break
			}
		}
	}
	fmt.Println(count)
}
