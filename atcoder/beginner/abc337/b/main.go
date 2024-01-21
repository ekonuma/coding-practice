package main

import(
	"fmt"
)

func main(){
	var s string
	fmt.Scan(&s)
	fmt.Println(extendedABC(s))
}

func extendedABC(s string) string{
	letters := []rune{'A', 'B', 'C'}
	current := 0

	for _, w := range s{
		if w != letters[current] {
			if current != 0 && w == letters[current-1]{
				return "No"
			}
			current++
		}
	}
	return "Yes"
}