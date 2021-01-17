package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//JSON İşlemleri

//Type nesneleri

//name struct'ını tanımlıyoruz
type Name struct {
	Family   string
	Personal string
}

//email struct
type Email struct {
	ID     int
	Kind   string
	Adress string
}

//ınterest struct
type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAdress(p *Person, i int) string {
	return p.Email[i].Adress
}

func GetPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarline() {
	fmt.Println("*******************")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

func saveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func main() {
	person := Person{
		ID:        9,
		FirstName: "omer",
		LastName:  "hodo",
		UserName:  "omerhodo",
		Gender:    "true",
		Name:      Name{Family: "hodolar", Personal: "omer"},
		Email: []Email{
			Email{ID: 1, Kind: "Work", Adress: "hodo@ff.cc"},
			Email{ID: 2, Kind: "Home", Adress: "hodo@ff.cc"},
		},

		Interest: []Interest{
			Interest{ID: 1, Name: "GO"},
			Interest{ID: 2, Name: "python"},
			Interest{ID: 3, Name: "ruby"},
		},
	}

	WriteMessage("reading operation started")
	WriteMessage("personal full name:")
	WriteStarline()
	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarline()

	WriteMessage("\n")
	WriteMessage("personal email with index")
	WriteStarline()
	resEmail := GetPersonEmailAdress(&person, 1)
	WriteMessage(resEmail)
	WriteStarline()

	WriteMessage("\n")
	WriteMessage("email object with index")
	WriteStarline()
	resEmail2 := GetPersonEmail(&person, 0)
	fmt.Println(resEmail2)
	WriteStarline()

	WriteMessage("reading operation ended")
	WriteMessage("\n")

	WriteMessage("writing operation started")
	saveJSON("person.json", person)
	WriteMessage("writing operation ended")
}
