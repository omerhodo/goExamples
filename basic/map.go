package main

import "fmt"

func main() {
	/*
		myMap := make(map[int]string)
		myMap[43] = "foo"
		myMap[12] = "buzz"

		fmt.Println(myMap)
		fmt.Println(myMap[43])
	*/

	states := make(map[string]string)
	states["ist"] = "istanbul"
	states["ank"] = "ankara"
	states["ant"] = "antalya"
	states["erz"] = "erzurum"
	fmt.Println(states)
	fmt.Println(states["ank"])
	fmt.Println(states["ank"], states["ant"])

	deger := states["ist"]
	fmt.Println(deger)

	delete(states, "erz")
	fmt.Println(states)

	states["izm"] = "izmir"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("key değeri %v value değeri %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for j := range states {
		keys[i] = j
		i++
	}
	fmt.Println(keys)

	for b := range keys {
		fmt.Println(states[keys[b]])
	}
}
