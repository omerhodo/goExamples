package main

import (
	"fmt"
)

func main() {

	//kullanıcı oluşturma v1

	user1 := &User{
		ID:        "1",
		FirstName: "ömer",
		LastName:  "hodo",
		UserName:  "omerhodo",
		Age:       27,
		Pay: &Payment{
			Salary: 4000,
			Bonus:  300,
		},
	}
	fmt.Println(user1)
	// fmt.Println(user1.Pay.Salary)
	// fmt.Println(user1.Pay.Bonus)
	// fmt.Println(user1.GetFullName())
	// fmt.Println(user1.GetUserName())
	// fmt.Println(user1.GetPayment())
	//fmt.Println("maaş:" + strconv.FormatFloat(user1.GetPayment(), -1, 64)) //hata var

	//kullanıcı oluşturma v2
	fmt.Println("\n")

	fmt.Println("kullanıcı oluşturma v2")

	user2 := newUser()
	user2.FirstName = "ömer"
	user2.LastName = "hodo"
	user2.UserName = "omerhodo"
	user2.Age = 27
	user2.Pay = &Payment{Salary: 900, Bonus: 1000}
	fmt.Println(user2.GetFullName())
	fmt.Println(user2.GetPayment())
}

type User struct {
	ID        string
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

func newUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

//ödeme yapısı
type Payment struct {
	Salary float64
	Bonus  float64
}

//ödemenin yapıcı metodu
func NewPayment() *Payment {
	p := new(Payment)
	return p
}

//metodlar

func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}
