package main

import "fmt"

func main() {
	str1 := "hodo"
	str2 := "veli"
	str3 := "mehemt"
	hayda, _ := fmt.Println(str1, str2, str3)
	//fmt.Println(hayda)
	fmt.Printf("haydayı yazacak: %v\n", hayda)

	i := 42
	fmt.Printf("floatı yazdırmak için: %.5f \n", float64(i))

	x := 10
	y := "hodo"
	z := true

	fmt.Printf("dataların tipleri X:%T, Y:%T, Z:%T 'dir \n", x, y, z)
}
