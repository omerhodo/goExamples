package main

import (
	"fmt"
	"time"
)

func main() {

	x := fmt.Println

	xTime := time.Date(1993, 10, 11, 23, 50, 10, 0, time.UTC)
	now := time.Now()
	x(now)

	x("---------------")

	x(now.Year())
	x(now.Month())
	x(now.Day())
	x(now.Hour())
	x(now.Minute())
	x(now.Second())
	x(now.Nanosecond())
	x(now.Location())
	x(now.Weekday())

	x("-----------------")

	//Tarihleri karşılaştırma
	x(xTime.Before(now)) //xtime'ın now dan önce mi olduğunu boolean değer olarak döner
	x(xTime.After(now))  //aynı şekilde sonra mı olduğunu
	x(xTime.Equal(now))  //eşitmi olduğunu döner.

	diff := now.Sub(xTime) //sub fonksiyonu tarihler arasındaki farkı bulur.
	x(diff)

	x(diff.Hours())
	x(diff.Minutes())
	x(diff.Seconds())
	x(diff.Nanoseconds())
}
