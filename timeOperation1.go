package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Şu anki unix zamanı değeri:%v\n", time.Now().Unix()) //unix zamanını ekrana yazdırmak için kullanıyoruz
	time.Sleep(2 * time.Second)                                      //programın çalışmasını bekletmek için kullanıyoruz.
	fmt.Printf("Şu anki unix zamanı değeri:%v\n", time.Now().Unix())

	currentTime := time.Now()                             //burda zamanın currentTime adındaki değişkene aktardık
	fmt.Printf(currentTime.Format("2006-01-12 15:04:05")) //burda ise currentTime.Format fonksiyonu kullanarak format fonksiyonu içine tarihi yazdırmak istediğimiz formatı yazdırarak ekrana yazılmasını sağladık
}
