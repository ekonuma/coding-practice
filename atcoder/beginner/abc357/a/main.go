package main

import "fmt"

func main() {
	var n, m int8
	fmt.Scan(&n, &m)

	h := make([]int8, n)

	for i := 0; i < int(n); i++ {
		fmt.Scan(&h[i])
	}

	r := 0
	for i := 0; i < int(n); i++ {
		m -= h[i]
		if m < 0 || i == int(n) {
			break
		}
		r++
	}
	fmt.Println(r)
}
