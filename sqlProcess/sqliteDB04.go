//Insert İşlemi (Kayıt Ekleme)

package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./sqliteDB01.db")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var (
		// ID        int
		UserName  string
		Email     string
		Password  string
		FirstName  string
		LastName  string
		BirthDate string
		IsActive  bool
	)



stmt,err:=db.Prepare("INSERT INTO users(UserName, Email, Password,FirstName,LastName,BirthDate,IsActive)VALUES(?,?,?,?,?,?,?)")
UserName="del"
Email="de@sd.c"
Password="def"
FirstName="al"
LastName="ri"
BirthDate="09.04.2000"
IsActive=true
res,errStmt:=stmt.Exec(UserName,Email,Password,FirstName,LastName,BirthDate,IsActive)
if errStmt!=nil{
	log.Fatal(errStmt)
}
fmt.Println(res.LastInsertId())


}