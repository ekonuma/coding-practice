package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, s int
	fmt.Scanf("%d", &n)
	line := make(map[int]int)
	search := -1
	result := ""
	
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &s)
		line[s] = i
	}

	for i := 0; i < n; i++ {
		next := line[search]
		if i == 0 {
			result = strconv.Itoa(next)
		}else{
			result += fmt.Sprint(" ", next)
		}
		search = next
	}
	fmt.Println(result)
}

