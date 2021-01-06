package main

import (
	"log"
	"os"
)

//byteları bir dosyaya yazmak
func main() {
	//demo1.txt dosyasını yazılabiir bir doya olarak aç

	file, err := os.OpenFile(
		"demo1.txt",
		os.O_WRONLY|os.O_CREATE,
		0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("bu dosyaya yazdık3!\n")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %d bytes. \n", byteWritten)

}
