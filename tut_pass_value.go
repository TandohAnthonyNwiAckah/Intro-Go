package main

import "fmt"

// func updateName(x string) {
// 	x = "John"
// }

func updateName(x string) string {
	x = "John"

	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func tutPassValue() {

	// group A types -> strings, ints, bools, floats, arrays, structs

	name := "Tony"

	// updateName(name)

	name = updateName(name)

	fmt.Println(name)

	// group B types -> slices, maps, functions
	// pointer wrapper values

	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)

}
