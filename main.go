package main

import "fmt"

var score = 99.5

func main() {

	// fmt.Println("Hello, World!")

	// Print
	// tutPrint()

	// Export hello package
	// result := hello.SayHello("Tony")
	// fmt.Println(result)

	// Arrays & Slices
	// tutArray()

	// Loops
	// tutLoop()

	// Functions
	// tutFunc()

	// Maps
	// tutMap()

	// Pass by value
	// Go is a pass by value language.
	// tutPassValue()

	// Pointers
	// tutPointer()

	// Structs

	mybill := newBill("Tony bill")

	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg pie", 8.95)
	mybill.addItem("toffee pudding", 4.95)
	mybill.addItem("coffee", 3.25)

	mybill.updateTip(10)

	// fmt.Println("Bill Name:", mybill)
	fmt.Println(mybill.format())

}
