package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	if strings.Index(s, "R") < strings.Index(s, "M") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
