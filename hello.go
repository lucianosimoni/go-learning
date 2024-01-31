package main

import "fmt"

func main() {
	helloName()
}

func helloWorld() {
	fmt.Println("Hello, World!")
}

func helloName() {
	fmt.Println("Please enter your name.")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi, %s! I'm Go!", name)
}