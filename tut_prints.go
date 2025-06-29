package main

import "fmt"

func tutPrint() {

	fmt.Print("Hello,")   // Print without newline
	fmt.Println("World!") // Print with newline

	age := 25
	name := "Tony"
	// fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted string), %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)    // %v = value in default format
	fmt.Printf("my age is %q    and my name is %q \n", age, name) // %q = quotes
	fmt.Printf("age is of type %T \n", age)                       // %T is the type
	fmt.Printf("you scored %f points! \n", 225.55)                // %f = float format
	fmt.Printf("you scored %0.1f points! \n", 225.55)             // %0.2f = float with 2 decimal points

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

}
