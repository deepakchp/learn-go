package main

import "fmt"

func main() {

	//variable example

	var name1 string = "Deepak Poilil"

	var name2 = "Deepak P"

	name3 := "Deepak"

	fmt.Println(name1, name2, name3)

	//Array example

	var ages1 [3]int = [3]int{1, 2, 3}

	var ages2 = [3]int{3, 4, 5}

	ages3 := [3]int{6, 7, 8}

	fmt.Println(ages1, ages2, ages3, len(ages1))

	names := [2]string{"prathamesh", "deepak"}

	fmt.Println(names)

	//Slices

	var employeeNo = []int{100, 200, 300}

	fmt.Println(employeeNo, len(employeeNo))

	employeeNo = append(employeeNo, 400)

	fmt.Println(employeeNo, len(employeeNo))

	// slice ranges

	employee := []string{"prathamesh", "deepak", "sandeep", "savitha"}

	range1 := employee[1:4]

	fmt.Println(range1)

}
