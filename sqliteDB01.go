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





	
/*
	createStatement := "'users'('ID' INTEGER PRIMARY KEY AUTOINCREMENT,'Username' varchar(45) NOT NULL,'Email' varchar(45) NOT NULL,'Password' varchar(45) NOT NULL,'FirstName' varchar(45) NOT NULL,'LastName' varchar(45) NOT NULL,'BirthDate' varchar(45) DEFAULT NULL,'IsActive' tinyint(1) DEFAULT NULL);"

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS" + createStatement)
	if err != nil {
		log.Fatal(err)
	}

	//veri ekleme işlemi
	res, err := db.Exec("INSERT INTO users(UserName, Email, Password, FirstName, LastName, BirthDate, IsActive) VALUES('hdil','hd@gmail.com','123456','harun','dilek','10.01.1993',1)")
	if err != nil {
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Inserted %d rows", rowCount)

	lastID, err1:=res.LastInsertId()
	if err1 != nil {
		log.Fatal(err1)
	}
log.Printf("Inserted ID: %d",lastID)

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

	//Eklenen kayıtları getir.
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&ID, &UserName, &Email, &Password, &FirsName, &LastName, &BirthDate, &IsActive)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("bulunan satır içeriği: %q ", strconv.Itoa(ID)+" "+UserName+" "+Email+" "+Password+" "+FirsName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
	}
	
*/


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

rows, errQ:= db.Query("SELECT * FROM users WHERE ID = ?", 2)
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
