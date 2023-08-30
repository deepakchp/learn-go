package main

import (
	"fmt"
	"sort"
	"strings"
)

func stndlibrary() {

	greetings := "Hello Go this is amazing"

	fmt.Println(strings.Index(greetings, "Go"))
	fmt.Println(strings.Contains(greetings, "Go"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Split(greetings, " "))

	ages := []int{10, 3, 7, 19, 4, 5}

	//sort.Ints(ages)

	fmt.Println(ages)

	index := sort.SearchInts(ages, 10)

	fmt.Println(index)
}
