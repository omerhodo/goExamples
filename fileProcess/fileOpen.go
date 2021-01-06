package main

import (
	"log"
	"os"
)

func main() {

	// file, err := os.Open("demo.txt")
	// defer file.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	//ilk parametre dosya adı, ikinci parametre dosya açılış amacı, üçüncü parametre dosya izinleri
	file, err := os.OpenFile("demo.txt", os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

}
