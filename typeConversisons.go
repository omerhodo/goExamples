package main

import (
	"fmt"
)

func main() {
	/*
		var myStrig string = "1"
		var x int = 10
		var y float32 = 2.0

		fmt.Println(myStrig, x, y)

		//stringten inte dönüştürme
		number, _ := strconv.Atoi(myStrig)
		fmt.Println(number)
		result := 2 + number
		fmt.Println(result)
		fmt.Println("Sonuç:" + strconv.Itoa(result))
	*/

	var i int = 55
	//var f1 float64 = i
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, u)
}
