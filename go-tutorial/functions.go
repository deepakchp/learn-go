package main

import (
	"fmt"
	"math"
	"strings"
)

func gofun() {
	saygreeting("Deepak")

	cloudnames := []string{"gcp", "azure", "aws", "oracle"}

	getcloudNames(cloudnames, saygreeting)

	area := circleArea(5.5)

	fmt.Println(area)

	firstname, lastname := getNames("Deepak Poilil")

	fmt.Println(firstname, lastname)
}

func saygreeting(n string) {
	fmt.Println("Hello ", n)
}

func getcloudNames(names []string, f func(n string)) {

	for _, val := range names {
		f(val)
	}
}

func circleArea(r float32) float32 {
	return math.Pi * r * r
}

func getNames(name string) (string, string) {

	n := strings.Split(name, " ")
	return n[0], n[1]
}
