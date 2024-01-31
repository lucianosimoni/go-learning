package main

import "fmt"

func main() {
	memoryPointer()
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

func memoryPointer() {
	// & = Is a variable that stores the Memory Address of another variable.
	number := 42

	pointer := &number

	fmt.Println(pointer)
}