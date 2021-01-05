package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2018, time.November, 10, 20, 0, 0, 0, time.UTC) //time package'inin Date fonksiyonunu kullanarak içine alması gereken parametreleri söylüyoruz. Date functionı belirli bir tarihi yazmak için kullanılır.
	fmt.Printf("çalışıyor: %s\n", t)                               //t'ye atadığımız değeri ekrana yazdırmayı söylüyoruz.

	fmt.Println("**************************")

	now := time.Now() //now functionı ile şu anki zamanı yazdırabiliriz.
	fmt.Printf("Mevcut zaman: %s\n", now)

	fmt.Println("**************************")

	fmt.Println("ay:", now.Month())              //now da üretilen şu anın tarihinin ayını ekrana yazdırdık
	fmt.Println("gün:", now.Day())               //aynın kaçıncı günü olduğunu ekrana yazdırdık
	fmt.Println("haftanın günü:", now.Weekday()) //haftanın günü
	fmt.Println("dakika:", now.Minute())
	fmt.Println("saniye:", now.Second()) //burda saniyeyi yazdırıp,
	time.Sleep(2 * time.Second)          //burda programın çalışmasını geciktiriyoruz,
	fmt.Println("saniye:", now.Second()) //burda yine aynı saniyeyi yazar çünkü biz now zamanını yukarıda oluşturmuştuk. burda yeniden bir zaman oluşmuyor.

	fmt.Println("**************************")

	//tarih ekleme
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("tomorrow date:%v, %v, %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year()) //now değişkenindeki tarihe bir gün eklemiş olduk.
	nextYearNow := now.AddDate(1, 0, 0)
	fmt.Printf("Next year now:%v, %v, %v, %v\n", nextYearNow.Weekday(), nextYearNow.Month(), nextYearNow.Day(), nextYearNow.Year()) //now değişkenindeki tarihe bir gün eklemiş olduk.

	//formatlı tarih yazma
	longFormat := "Moday,january 1,2000"                     //tarihi nasıl yazdırmak istediğimiz belirtmek için kullanacağız.
	fmt.Println("tomorrow is ", tomorrow.Format(longFormat)) //format functionı ile tarihi nasıl yazdırmak istediğimiz söyledik. içine parametre olarak daha önceden belirlediğimiz formattaki değişkeni gönderdik.
	shortFormat := "1/2/2000"
	fmt.Println("tomorrow is ", tomorrow.Format(shortFormat))
}
