package main

import (
	"fmt"
)

type strFloatMap map[string]float64

func (m strFloatMap) output() {
	fmt.Println(m)
}

func main() {

	websites := map[string]string{
		"google": "www.google.com",
		"aws":    "www.aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["aws"])
	websites["azure"] = "www.azure.com"

	fmt.Println(websites)
	//to delete key
	delete(websites, "azure")
	fmt.Println(websites)

	/**
	  we can use make builtin func to create  array or map
	  2: initial size of the array
	  5: It is not mandatory, it defines capacity of the array
	*/
	userNames := make([]string, 2, 5)
	//userNames := [] string{}
	// commented code throws exception if you want to get specific index userNames[1]
	// because it is not created yet.
	userNames[0] = "Julie"

	userNames = append(userNames, "Max")

	fmt.Println(userNames)

	// 3 defines the capacity of the map
	// if you exceed the capacity go tries to reallocate your map in to memory
	courseRatings := make(map[string]float64, 3)
	courseRatings["go"] = 4.9
	courseRatings["python"] = 4.1
	courseRatings["react"] = 4.7
	fmt.Println(courseRatings)

	courseRatings2 := make(strFloatMap, 3)
	courseRatings2["go"] = 4.9
	courseRatings2["python"] = 4.1
	courseRatings2["react"] = 4.7
	courseRatings2.output()

	for index, value := range userNames {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}

	for key, value := range courseRatings {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}
}
