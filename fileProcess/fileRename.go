//dosyanın konumunu ve adını değiştirme işlemi.
package main

import (
	"log"
	"os"
)

func main(){
	orginalPath:="demo.txt" //dosyanın burda olup olmadığını kontrol edin
	newPath:="./moved/test.txt" //bu adda bir dosya olup olmadığını kontrol edin
	err:=os.Rename(orginalPath,newPath)
	if err !=nil{
		log.Fatal(err)
	}
}