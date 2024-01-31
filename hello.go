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

/* Quickly explains how the & and * works in Go */
func memoryPointer() {
	// & = Is a variable that stores the Memory Address of another variable.
	// Declare a variable 'num' and assign a value
    num := 42

    // Declare a pointer variable 'ptr' of type int, and assign the memory address of 'num' to it
    ptr := &num

    fmt.Println("Value of 'num':", num)
    fmt.Println("Memory address of 'num':", &num)
    fmt.Println("Value pointed to by 'ptr':", *ptr) // Dereferencing the pointer to access the value

    // Modify the value of 'num' indirectly through the pointer
    *ptr = 100
    fmt.Println("Modified value of 'num':", num)
}