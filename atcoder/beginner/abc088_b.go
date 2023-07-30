package main

import (
	"fmt"
	"sort"
)

func main()  {
	var alice, bob,n int
	var pointNum []int

	fmt.Scan(&n)
	pointNum = make([]int, n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&pointNum[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(pointNum)))

	for i:=0; i<n; i++{
		if i%2 == 0{
			alice+=pointNum[i]
		}else{
			bob+=pointNum[i]
		}
	}
	alice -= bob
	fmt.Println(alice)	
}