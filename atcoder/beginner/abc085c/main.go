package main

import "fmt"

func main()  {
	var n, y int
	fmt.Scan(&n)
	fmt.Scan(&y)
	fmt.Println(otoshidama(n, y))
}

func otoshidama(n int, y int) string{
	for itiman := 0; itiman <= n; itiman++{
		for gosen := 0; gosen <= n-itiman; gosen++{
			sen := n - itiman - gosen
			if itiman * 10000 + gosen * 5000 + sen * 1000 == y{
				return fmt.Sprintln(itiman, gosen, sen)
			}
		}	
	}
	return fmt.Sprintln("-1", "-1", "-1") 
}