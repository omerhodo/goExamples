package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// func NewHuman() *Human {
// 	h := new(Human)
// 	return h
// }

func NewHumanWithParams(firstName, lastName string, age int) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = lastName
	h.Age = age
	return h
}

func main() {

	/*

		//Nesne Örneği Oluştur
		//v1...
		x := Human{FirstName: "Omer"}
		fmt.Println(
		//v2
		x := new(Human)
		x.FirstName = "Ömer"
		fmt.Println(x.FirstNam
		//yapıcı metod kullanı
		x := NewHuman()
		x.Age = 28
		fmt.Println(x.Age)

		//parametreli kullanım
		x := NewHumanWithParams("ömer", "hodo", 27)
		fmt.Println(x.FirstName, x.LastName, x.Age)

		//
		x := NewHumanWithParams("ömer", "hodo", 27)
		var data = x.FirstName + " " + x.LastName + "/" + strconv.Itoa(x.Age)
		fmt.Println(data)
	*/
	v := Vertex{1, 2}
	p := &v
	p.x = 1e9
	fmt.Println(v)

}
