package main

import "fmt"

func main() {
	/*
		var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

		for i, v := range pow {
			fmt.Printf("2**%d=%d\n", i, v)
		}

		//array
		a := [...]string{"a", "b", "c", "d", "f"}
		for i := range a {
			fmt.Println("array item", i, "is", a[i])
		}
		//map

		baskentler := map[string]string{"Turkey": "Ankara", "France": "Paris", "Italy": "Roma", "Japan": "Tokyo", "Almanya": "Berlin"}

		for key := range baskentler {
			fmt.Println("ülke:", key, "başkent:", baskentler[key])
		}
	*/

	//map

	baskentler := map[string]string{"Turkey": "Ankara", "France": "Paris", "Italy": "Roma", "Japan": "Tokyo", "Almanya": "Berlin"}

	for key, value := range baskentler {
		fmt.Println("ülke:", key, "başkent:", value)
	}
}
