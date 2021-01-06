//sql ile single row öneği

package main

import (
	"database/sql"
	"log"
	// "strconv"
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


/*
err=db.QueryRow("SELECT*FROM users limit 1").Scan(&ID, &UserName, &Email, &Password, &FirsName, &LastName, &BirthDate, &IsActive)
if err !=nil{
	if err==sql.ErrNoRows{

	}else{
		log.Fatal(err)
	}
}
log.Printf("bulunan satır içeriği: %q ", strconv.Itoa(ID)+" "+UserName+" "+Email+" "+Password+" "+FirsName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
*/

//Multiple Active Result Set: MARS
//_, err:=db.Exec("DELETE FROM xTable1; DELETE FROM xTable2")


//Preparing Queries
/*
stmt, errQ:=db.Prepare("SELECT * FROM users WHERE ID=?")
if errQ != nil{
	log.Fatal(errQ)
}
defer stmt.Close()
rows, errX:=stmt.Query(3)
if errX!=nil{
	log.Fatal(errX)
}
defer rows.Close()
for rows.Next(){
	scanErr:=rows.Scan(&ID, &UserName, &Email, &Password, &FirsName, &LastName, &BirthDate, &IsActive)
	if scanErr!=nil{
		log.Fatal(scanErr)
	}
	log.Printf("bulunan satır içeriği: %q ", strconv.Itoa(ID)+" "+UserName+" "+Email+" "+Password+" "+FirsName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
}
*/


//Preparing Query - Single Rows
/*
	stmt, errQ:= db.Prepare("SELECT*FROM users WHERE ID = ?")
	if errQ!=nil{
		log.Fatal(errQ)
	}
	errX:=stmt.QueryRow(1).Scan(&ID, &UserName, &Email, &Password, &FirsName, &LastName, &BirthDate, &IsActive)
	if errX !=nil{
		log.Fatal(errX)
	}
	fmt.Println(FirsName+" "+LastName)

*/



//Modifying Data
/*
// burda bir hata var

res, _ := db.Exec("DELETE FROM users LIMIT 1")
rowCount, _:=res.RowsAffected()
fmt.Println(rowCount)
lastID,_:=res.LastInsertId()
fmt.Println(lastID)
*/



}
