package main

import "fmt"

func main() {

	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "f50"
	ferr.Color = "Red"
	ferr.Speed = 420
	ferr.Price = 40000999000
	ferr.Special = true
	fmt.Println(ferr.Information())
}

//interface
type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

//structs
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}

//Ferrari

type Ferrari struct {
	Car               //composition metodu ile yukarıdaki structı car struct'ını dahil ettik.
	SpecialProduction //composition ile bnu da dahil ettik.
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (x *Ferrari) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color :" + x.Color
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special: " + add
	}
	return ret
}

//Lamborgini

type Lamborgini struct {
	Car               //composition metodu ile yukarıdaki structı car struct'ını dahil ettik.
	SpecialProduction //composition ile bnu da dahil ettik.
}

func (_ Lamborgini) Run() bool {
	return true
}

func (_ Lamborgini) Stop() bool {
	return true
}

func (x *Lamborgini) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color :" + x.Color
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special: " + add
	}
	return ret
}

//Mercedes

type Mercedes struct {
	Car               //composition metodu ile yukarıdaki structı car struct'ını dahil ettik.
	SpecialProduction //composition ile bnu da dahil ettik.
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (x *Mercedes) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color :" + x.Color
	return ret
}
