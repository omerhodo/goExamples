package main

import (
	"runtime"
	"time"
)

/*
Bu projede iki tane go routine ve bir tane de main routine olmak üzere üç tane routine oluşturuldu.
İki go routine functionların çalışması için kullanıldı ve main routine genel routine'miz.
*/

func main() {
	runtime.GOMAXPROCS(8)
	ch := make(chan string) //string tipinde bir channel tanımladık
	go xFunc(ch)            //xFunc function'ı ile veriyi hesaplıyoruz.
	go printer(ch)          //burda da hesaplanan veriyi ekrana yazdırıyoruz.
	time.Sleep(100 * time.Millisecond)

}

func xFunc(ch chan string) { //dışarıdan channel alan bir function oluşturduk
	for l := byte('a'); l <= byte('z'); l++ { //burda l adındaki değişkeni a'nın byte türündeki değerinden, z'nin byte türündeki değerine kadar arttırdık.
		ch <- string(l) //herbir byte türündeki değişkeni sırayla ch'ye aldık
	}
}

func printer(ch chan string) { //yine bir channel alan function oluşturduk
	for {
		println(<-ch) //ch'yi ekrana yazdırıyoruz.
	}
}
