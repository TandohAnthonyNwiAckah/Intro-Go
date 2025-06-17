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

	name := "Tony"

	// updateUser(name)
	// fmt.Println(name)

	// & gets the memory address of the value (pointer)
	// fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	m := &name
	// fmt.Println("memory address:", m)
	// fmt.Println("value at memory address:", *m)

	fmt.Println(name)
	updateUser(m) // using pointer as arg, can pass &name directly instead of "m" as well
	fmt.Println(name)

}

/*

|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "Tony"  | p0x001  |
|---------|---------|

*/
