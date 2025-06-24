package main

import "fmt"

// func updateUser(n string) {
// 	// fmt.Println(n)
// 	n = "John"
// 	// fmt.Println(n)

// }

func updateUser(n *string) {
	*n = "John"
}

func tutPointer() {

	// name := "Tony"

	// updateUser(name)
	// fmt.Println(name)

	// & gets the memory address of the value (pointer)
	// fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	// m := &name
	// fmt.Println("memory address:", m)
	// fmt.Println("value at memory address:", *m)

	// fmt.Println(name)
	// updateUser(m) // using pointer as arg, can pass &name directly instead of "m" as well
	// fmt.Println(name)

	// TESTING POINTERS

	limit := 25

	points := &limit // points is a pointer to limit

	fmt.Println("POINTS:", points)
	fmt.Println("POINTS POINTER:", *points)

	*points = *points + 5 // incrementing the value at the memory address by 5

	fmt.Println("LIMIT:", limit)

}

/*

|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "Tony"  | p0x001  |
|---------|---------|

*/
