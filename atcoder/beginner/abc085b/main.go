package main

import(
	"fmt"
)

func main()  {
	var n int8
	fmt.Scan(&n)
	mochiSize := make([]int8, n)
	
	for i := 0 ; int8(i) < n; i++ {
		fmt.Scan(&mochiSize[i])
	}

	fmt.Println(kagamimochi(mochiSize...))
}

func kagamimochi(size...int8) int{
	mochis := map[int8]struct{}{}
	for _, mochi  := range size {
		mochis[int8(mochi)]= struct{}{}
	}
	return len(mochis)
}

func execTestCases(){
	fmt.Println(kagamimochi(4, 10, 8, 8, 6))
	fmt.Println(kagamimochi(3, 15, 15, 15)) 	
	fmt.Println(kagamimochi(7, 50, 30, 50, 100, 50, 80, 30))
}