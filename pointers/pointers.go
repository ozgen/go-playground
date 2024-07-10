package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age // to get address value of a poınter use &

	fmt.Println("Age: ", *agePointer) // to get the real value of from the pointer use *

	getAdultYears(agePointer)
	fmt.Println("Adult years: ", *agePointer)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
