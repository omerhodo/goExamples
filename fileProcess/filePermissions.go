//okuma ve yazma izinleri kontrolü yapmak

package main

import (
	"log"
	"os"
)

func main() {
	// //yazma izinleri
	// file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)
	// if err != nil {
	// 	if os.IsPermission(err) {
	// 		log.Println("hata:yazma izni reddedildi.")
	// 	}
	// }
	// file.Close()

	//okuma izinleri
	file, err := os.OpenFile("demo.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("hata:okuma izni reddedildi.")
		}
	}
	file.Close()
}
