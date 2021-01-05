package main

import (
	"fmt"
	"time"
)

func main() {
	//konsol saat tarih işlemleri

	//DATE(METODU İLE TARİH VERİMİZİ OLUŞTURUYORUZ.)
	t := time.Date(2016, time.November, 10, 20, 0, 0, 0, time.UTC)

	//t adıyla tarih tipinde oluşturulan veriyi string tipined ekrana yazıyoruz.
	fmt.Printf("go launch at %s\n", t)

	fmt.Println("-----------------")

	//time.Now() ile şu anın zamanını elde edebiliyoruz
	now := time.Now()

	//elde edilen zaman bilgisini ekrana böyle yazdırıyoruz
	fmt.Printf("the time now is %s\n", now)

	fmt.Println("-----------------")

	//ilk başta oluşturduğumuz t adındaki zaman bilgisini ay,gün ve haftanın günü olarak ekrana yazdırıyoruz.
	fmt.Println("the month is", t.Month())
	fmt.Println("the day is", t.Day())
	fmt.Println("the weekday is", t.Weekday())

	fmt.Println("-----------------")

	//tarihe bir gün ekle

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("tomorrow is %v, %v, %v, %v \n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())

	fmt.Println("-----------------")

	longFormat := "monday, january 2, 2006"
	fmt.Println("tomorrow is", tomorrow.Format(longFormat))

	shortFormat := "1/2/2006"
	fmt.Println("tomorrow is", tomorrow.Format(shortFormat))

}
