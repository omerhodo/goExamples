package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Println("sumın değeri: ", sum)
		time.Sleep(300 * time.Millisecond)
	}
}
