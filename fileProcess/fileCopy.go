package main

import (
	"io"
	"log"
	"os"
)

//dosya kopyalama
func main() {

	orginalFile, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer orginalFile.Close()

	//Yeni bir dosya oluşturma
	newFile, err := os.Create("./folder/demoCopy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//byte'ları kaynaktan hedefe kopyala
	bytesWritten, err := io.Copy(newFile, orginalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes", bytesWritten)

	//dosya içeriğini işle
	//belleği diske boşalt
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
