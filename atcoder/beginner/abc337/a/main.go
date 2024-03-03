package main

import (
	"fmt"
)

func main() {
	var n, x, y, takahashi, aoki int

	fmt.Scan(&n)

	for i := 0; i < (n); i++ {
		fmt.Scanln(&x, &y)
		takahashi += x
		aoki += y
	}

	if takahashi == aoki {
		fmt.Println("Draw")
	} else if takahashi > aoki {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
