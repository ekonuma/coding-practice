package main

import (
	"fmt"
)

func main() {
	var n, a, b, ans int
	fmt.Scan(&n, &a, &b)

	for i := a; i <= n; i++ {
		var temp int
		for j := i; j > 0; j /= 10 {
			temp += j % 10
		}
		if temp >= a && temp <= b {
			ans += i
		}
	}
	fmt.Println(ans)
}
