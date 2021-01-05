package main

import (
	"fmt"
	"strconv"
)

func main() {

	student := &Student{
		ID:        "1",
		ogrNo:     43,
		FirstName: "ömer",
		LastName:  "hodo",
		FullName:  "ömerhodo",
		Age:       27,
		ort: &Ortalama{
			dortluk: 3.57,
			yuzluk:  83,
		},
	}

	fmt.Println(student)
	fmt.Println(student.getFullName())
	fmt.Println(student.getOgrNo())
	fmt.Println(student.sumOrtalama())

	fmt.Println("\n")
	fmt.Println("\n")

	fmt.Println("connection")
	fmt.Println(student.herhangi())
}

type Student struct {
	ID        string
	ogrNo     int
	FirstName string
	LastName  string
	FullName  string
	Age       int
	ort       *Ortalama
}

//ortalama girme
type Ortalama struct {
	dortluk float64
	yuzluk  float64
}

//yeni öğrenci oluşturma
func newStudent() *Student {
	s := new(Student)
	s.ort = newOrtalama()
	return s
}

//ortalama belirleme
func newOrtalama() *Ortalama {
	o := new(Ortalama)
	return o
}

//metodlar
func (s Student) getFullName() string {
	return s.FirstName + " " + s.LastName
}

//ogrNo bul

func (s Student) getOgrNo() int {
	return s.ogrNo
}

//ortalama toplama
func (s Student) sumOrtalama() float64 {
	return s.ort.dortluk + s.ort.yuzluk
}

//herhangi bir şey yazdır
func (s Student) herhangi() string {
	return s.FirstName + " " + s.LastName + " " + s.FullName + " " + strconv.Itoa(s.ogrNo)
}
