package main

import(
	"fmt"
)

func main(){
	var a, b int16
    var seisuu string
	fmt.Scanf("%d %d", &a, &b);

	if(a * b % 2 == 0){
       seisuu = "Even"
	}else {
	   seisuu = "Odd"
	}
	fmt.Println(seisuu)
}

