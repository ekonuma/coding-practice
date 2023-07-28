package main

import "fmt"

func main() {
	var a, b, c, x, ans int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for h := 0; h <= c; h++ {
				if i*500+j*100+h*50 == x {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
