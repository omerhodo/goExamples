package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	//geçici klasör oluştur
	tempDirPath, err := ioutil.TempDir("", "geçiciKlasör")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("geçici klasör oluşturduk:", tempDirPath)

	//geçici dosya oluştur.
	tempFile, err := ioutil.TempFile(tempDirPath, "geciciDosya.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("geçici dosya oluşturduk: ", tempFile.Name())

	//close file
	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	//silme
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s dosyası silindi", tempFile.Name())
	}

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}

}
