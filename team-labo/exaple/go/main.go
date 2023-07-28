package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	MultipleOf7Add()
	MultipleOf7OneByOne()
}

func MultipleOf7Add(){
	num := 7
	mult7 := 0
	for num <= 7777777  {
		mult7 += strings.Count(strconv.Itoa(num) ,"7")
		num+=7
	}
	fmt.Println("result", mult7)
}

func MultipleOf7OneByOne(){
	num := 7
	mult7 := 0
	for num <= 7777777  {
		if num % 7 == 0{
			mult7 += strings.Count(strconv.Itoa(num) ,"7")
		}
		num++
	}
	fmt.Println("result", mult7)
}
