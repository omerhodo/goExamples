package main

import (
	"fmt"
	"os"
)

/*

func main() {
	file, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println("burda hata var")
	} else {
		fmt.Println(file.Name)
	}
}

*/

/*
func main() {
	myError := errors.New("bu bir hata")
	fmt.Println(myError)
}

*/

func main() {
	_, err := os.Open("abc.rar")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
