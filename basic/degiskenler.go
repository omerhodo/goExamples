package main

import "fmt"

var g, h, j = 43, "esey", 54 //burda var ile değişken tanımlamak mümkün. ancak := ile tanımlayıp daha sonra kullanılamaz.
const pi = 3.14              //sabit atama yapmak mümkün. aşağıdaki alanda değiştirilemez.

func main() {

	/*
		var message string
		message = "merhaba go"
		var message = "mert hakan"
		var a, b, c int = 3, 4, 5

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)

		var a, b, c = 1, true, "merhaba"
		var endegisik = "hayda"
		fmt.Println(a, b, c, endegisik)

		v := 55
		fmt.Println(v)
		a, b := "hodo", 54
		fmt.Println(a, b)
		message := "hodolar"
		fmt.Println(message)

	*/
	fmt.Println(g, h, j)

	var (
		degisken1 = "omer"
		degisken2 = "hodo"
	)
	fmt.Println(degisken1, degisken2)
}
