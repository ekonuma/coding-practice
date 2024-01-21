package main

import (
	"fmt"
	
	"strings"
)

func main() {
	var s string	
	fmt.Scan(&s)
	fmt.Println(dayDream(s))
}

func dayDream(s string) string {
	ddee := []string{"dream", "dreamer", "erase", "eraser"}
	size := len(s)
	currentSize := size
	for {
		for i := 0; i < 4; i++ {
			if strings.HasSuffix(s[:size], ddee[i]) {
				size -= len(ddee[i])
				break
			}
		}
		if size == 0 {
			return "YES"
		}
		
		if currentSize == size{
			return "NO"
		}
		currentSize = size
	}
}