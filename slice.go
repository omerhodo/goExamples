package main

import "fmt"

func main() {
	/*
		örn1
		primes := []int{1, 2, 3, 4, 5} //dizimizi burda tanımladık

		var s []int = primes[1:4] //slice tanımlamasını burda yaptık. primes[1:4]'diyerek dizinin 1.indisinden 4.indisine kadar olan bölümü aldık.
		fmt.Println(s)
		fmt.Println(primes)

		//örn2
		myArray := [...]int{45, 23, 12}
		mySlice := myArray[:]
		fmt.Println(mySlice)
		mySlice[0] = 11      //slice üzerinden arrayin verisini değiştirdik.
		fmt.Println(mySlice) //slice'ın üzerinden arrayi değiştirdik
		fmt.Println(myArray) //arrayin verisi artık değişti

		//örn3
		var colors = []string{"red", "green", "blue"}
		fmt.Println(colors)
		fmt.Println(colors)
		colors = append(colors, "purple") //colors slice'ına purple'ı eklemek için append fonksiyonunu kullanırız.
		colors = append(colors[1:len(colors)]) //append fonksiyonuna nerden nereye kadar güncelleme yapacağını söyleriz. burdada color slice'ının 1.indisli elemanından başlayıp colors slice'ının len fonksiyonuyla uzunluğunu alarak o uzunluğa kadar devam etmesini söyledik.
		fmt.Println(colors)
		colors = append(colors[:len(colors)-1]) //bu sefer colors slice'ının sondan bir önceki verisine kadar gitmemizi yani son elemanı almamamızı söyledik.
		fmt.Println(colors)
	*/

	numbers := make([]int, 5, 5)
	numbers[0] = 123
	numbers[1] = 451
	numbers[2] = 34
	numbers[3] = 45
	numbers[4] = 543

	fmt.Println(numbers)
	numbers = append(numbers, 142)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))

}
