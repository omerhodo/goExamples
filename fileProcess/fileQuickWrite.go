package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

//hızlı dosya yazma işlemi
func main() {
	err := ioutil.WriteFile("demo2.txt", []byte("bu veriyi demo2ye yazacak"), 0666)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("işlem başarılı")
	}
}
