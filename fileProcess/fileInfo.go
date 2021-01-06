//Get File İnfo

package main

import (
	"fmt"
	"log"
	"os"
)

var (
	// fileInfo *os.fileInfo
	err error
)

func main(){
	//dosya bilgisini hata yoksa fileInfo'ya döner. hata varsa err'de tutulur
	fileInfo, err :=os.Stat("demo.txt")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("file name:",fileInfo.Name())
	fmt.Println("size in bytes:",fileInfo.Size())
	fmt.Println("permissions:",fileInfo.Mode())
	fmt.Println("last modified:",fileInfo.ModTime())
	fmt.Println("Is Dictionary:",fileInfo.IsDir())
	fmt.Println("System Interface Type:",fileInfo.Sys())
	
}