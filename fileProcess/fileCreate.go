//File Creating Operation

package main


import (
	"log"
	"os"
)

var (
	newFile *os.File
	err error
)

func main() {
newFile, err=os.Create("olusanDosyaAdi.txt")
if err!=nil{
	log.Fatal(err)
}
}