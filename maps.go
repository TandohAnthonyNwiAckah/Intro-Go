package main

import "fmt"

func tutMap() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "doe",
		984759373: "jane",
		845775485: "john",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[845775485])

	phonebook[984759373] = "tony"
	fmt.Println(phonebook)

	phonebook[647583927] = "mary"
	fmt.Println(phonebook)
}
