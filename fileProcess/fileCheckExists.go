//dosyanın var olup olmadığını sorgulama

package main

import (
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("file does not exist")
		}
	}
	log.Println("file does exist. file information:")
	log.Println(fileInfo)
}
