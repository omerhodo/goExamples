package main

import "fmt"

func main() {
	/*
	   var a [10]string
	   a[0] = "1"
	   a[1] = "2"
	   a[2] = "3"
	   a[3] = "100"
	   a[4] = "1001"
	   a[5] = "133"
	   fmt.Println(a)

	   fmt.Println(a[4], a[5])

	   primes := [6]int{2, 3, 5, 7, 8}
	   fmt.Println(primes)
	   fmt.Println(a[1] + a[4])

	*/

	/*


		var numbers = [5]int{1, 23, 4, 67, 4}
		fmt.Println(numbers)
		fmt.Println("listedeki toplam sayı adedi:", len(numbers))
	*/

	/*

		myArray := [...]int{4, 54, 453, 23, 234, 12}
		fmt.Println(myArray)
	*/
	var cars [4]string
	cars[0] = "Tesla"
	cars[1] = "Audi"
	cars[2] = "Mercedes"
	cars[3] = "Porcshe"

	/*

		//yöntem1
		i := 0
		for i <= len(cars)-1 {
			fmt.Println(cars[i])
			i++
		}
	*/

	/*

		//yöntem2
		for i := 0; i < len(cars); i++ {
			fmt.Println(cars[i])
		}
	*/

	/*
		//yöntem3

		for i, value := range cars {
			fmt.Println("i", i, "&value", value)
		}
	*/

	/*
		//yöntem4
	*/

	for _, value := range cars {
		fmt.Println("arabalar:", value)
	}
}
