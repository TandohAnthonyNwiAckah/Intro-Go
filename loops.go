package main

import "fmt"

func tutLoop() {

	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++ // infinite loop without this
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"tony", "jane", "joe", "mary"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, val := range names {
		fmt.Printf("the value at pos %v is %v \n", index, val)
		val = ""
	}

	for _, val := range names {
		fmt.Print(val, ",")
		val = "new string"
	}

	// changing val in a loop does not mutate the original slice
	fmt.Println(names)

	// booleans & conditionals

	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 40")
	}

	anames := []string{"tony", "mary", "joe", "doe", "jane"}

	for index, val := range anames {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, val)
	}
}
