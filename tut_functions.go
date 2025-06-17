package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func tutFunc() {

	sayGreeting("John Doe")

	cycleNames([]string{"cloud", "barret", "tifa"}, sayGreeting)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 area is %0.3f & circle 2 area is %0.3f \n", a1, a2)

	// Multiple Return Values

	fn1, sn1 := getInitials("john doe")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("cloud strife")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("tony")
	fmt.Println(fn3, sn3)

}
