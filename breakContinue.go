package main

func main() {
	/*

		i := 0
		for {
			if i >= 3 {
				break
			}
			println("i'nin değeri:", i)
			i++
		}
	*/

	for i := 0; i < 7; i++ {
		if i == 3 {
			continue
		} else if i == 5 {
			break
		}
		println("Çıktı:", i)
	}
	println("işlem devam...")
}
