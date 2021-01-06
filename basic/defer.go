package main

import "fmt"

var isConnected bool = false

func main() {
	fmt.Printf("connection open %v\n", isConnected)
	databaseProcessing()
	fmt.Printf("connection open %v\n", isConnected)

}

func databaseProcessing() {
	connect()
	fmt.Println("Deferring disconnect!")
	defer disconnect()
	fmt.Printf("connection open %v\n", isConnected)
	fmt.Println("Doing something")

}
func connect() {
	isConnected = true
	fmt.Println("Connected to database!")
}

func disconnect() {
	isConnected = false
	fmt.Println("Disconnected!")

}

func exit() {
	println("good bye!!")
}
