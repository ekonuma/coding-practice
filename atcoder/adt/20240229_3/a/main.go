package main

import "fmt"

func main() {
	var x1, y1, x2, y2, x3, y3 int8
	fmt.Scanf("%d %d", &x1, &y1)
	fmt.Scanf("%d %d", &x2, &y2)
	fmt.Scanf("%d %d", &x3, &y3)
	fmt.Printf("%d %d", x1^x2^x3, y1^y2^y3)
}
