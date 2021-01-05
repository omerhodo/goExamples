package main

import "fmt"

func main() {
	var score float64
	fmt.Println("notu girin:")
	fmt.Scanf("%v", &score)

	switch {
	case score <= 59:
		println("f")

	case score <= 69:
		println("d")

	case score <= 79:
		println("c")

	case score <= 89:
		println("b")

	case score <= 100:
		println("a")
	default:
		println("0ile 100 arası bir sayı girin!")
	}
}
