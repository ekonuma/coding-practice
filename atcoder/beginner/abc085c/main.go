package main

import "fmt"

func main()  {
	var n,y, sen, gosen, itiman, total int
	fmt.Scan(&n)
	fmt.Scan(&y)

	(10 * i) +  (5 * g) + s = y
	total = y

	for(y % 1000 == 0) {
		y -= y/1000 
		sen+= y/1000
	 }
	 if sen > n{
		y = total
		for(y % 5000 == 0) {
			y -=  y/5000 
			gosen+= y/5000
		 } 
		 if gosen > n{
			y = total
			for(y %10000 == 0) {
				y -=  y/10000
				itiman+= y/10000
			 } 
			 for(y % 5000 == 0) {
				y -=  y/5000
				gosen+= y/5000
			 }  
			 for(y % 1000 == 0) {
				y -=  y/1000
				sen+= y/1000
			 }  
		 }
	 }
	 for(y % 5000 == 0) {
		y -=  y/5000
		gosen+= y/5000
	 }
	 fmt.Println(itiman, gosen, sen)
}