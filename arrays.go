package main

import (
	"fmt"
	"sort"
	"strings"
)

func tutArray() {

	// var ages [3]int = [3]int{20, 25, 30}

	var ages = [3]int{20, 25, 30}

	names := [4]string{"tony", "john", "jane", "thomas"}

	names[1] = "tandoh"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:4] // doesn't include pos 4 element
	rangeTwo := names[2:]  //includes the last element
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	rangeOne = append(rangeOne, "yaba")
	fmt.Println(rangeOne)

	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "there")) // true
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "fr"))
	fmt.Println(strings.Split(greeting, " "))

	mages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(mages)
	fmt.Println(mages)

	index := sort.SearchInts(mages, 30)
	fmt.Println(index)

	ames := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(ames)
	fmt.Println(ames)

	fmt.Println(sort.SearchStrings(ames, "bowser"))

}
