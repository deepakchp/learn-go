package main

import (
	"fmt"
)

func loop() {

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	names := []string{"Deepak", "Prathamesh", "Sandeep", "Savitha"}

	for k := 0; k < len(names); k++ {
		fmt.Println(names[k])
	}

	for index, value := range names {
		fmt.Println(index, value)
	}
}
