//sql ile database'deki kayıtları getirme

package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./sqliteDB01.db")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close() //defer özelliği ile bu scope'un en son işlemi olduğunu söylendi.

	var (
		ID        int
		UserName  string
		Email     string
		Password  string
		FirsName  string
		LastName  string
		BirthDate string
		IsActive  bool
	)

rows, errQ:= db.Query("SELECT * FROM users WHERE ID = ?",3)
if errQ != nil {
	log.Fatal(errQ)
}
defer rows.Close()
for rows.Next(){
	err = rows.Scan(&ID, &UserName, &Email, &Password, &FirsName, &LastName, &BirthDate, &IsActive)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bulunan satır içeriği: %q ", strconv.Itoa(ID)+" "+UserName+" "+Email+" "+Password+" "+FirsName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))

}
errQ=rows.Err()
if errQ !=nil{
	log.Fatal(errQ)
}

	
	
}
