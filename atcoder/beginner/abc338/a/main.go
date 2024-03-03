package main

import (
	"fmt"
	"strconv"
)

func main(){
	var s string;
	fmt.Scan(&s);
	fmt.Println(capitalized(s))	
}

func capitalized(s string) string{
	for i, c := range s{
		char, _ := strconv.Atoi(fmt.Sprintf("%d", c))
		if i == 0 && char << 1 > 180{
			return "No"
		}else if i > 0 && char << 1 < 181{
			return "No"
		}
	}
	return "Yes"
}