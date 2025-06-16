package main

import "fmt"

func main() {

	// fmt.Println("Hello, World!")

	// strings
	// var nameOne string = "Alice"
	// fmt.Println("Name One:", nameOne)

	// var nameTwo = "Bob"
	// fmt.Println("Name Two:", nameTwo)

	// var nameThree string
	// fmt.Println("Name Three:", nameThree)

	// nameFour := "Charlie" // short variable declaration. only in functions
	// fmt.Println("Name Four:", nameFour)

	// ints
	// var ageOne int = 30
	// fmt.Println("Age One:", ageOne)
	// var ageTwo = 25
	// fmt.Println("Age Two:", ageTwo)
	// var ageThree int
	// fmt.Println("Age Three:", ageThree)
	// ageFour := 40 // short variable declaration. only in functions
	// fmt.Println("Age Four:", ageFour)

	// // bits & memory
	// var bitOne int8 = 25 // 8 bits = 1 byte
	// fmt.Println("Bit One:", bitOne)

	// var numOne int8 = -128
	// fmt.Println("Num One:", numOne)

	// var numTwo uint8 = 255
	// fmt.Println("Num Two:", numTwo)

	// float
	// var flOne float32 = 24.139
	// fmt.Println("Float One:", flOne)
	// var flTwo float64 = 2884292499249294
	// fmt.Println("Float Two:", flTwo)
	// var flThree float64
	// fmt.Println("Float Three:", flThree)

	// Print
	// fmt.Print("Hello,")   // Print without newline
	// fmt.Println("World!") // Print with newline

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
