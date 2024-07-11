package main

import "fmt"

func main2() {
	var productNames [4]string = [4]string{"book"}
	prices := [4]float64{1.2, 3.0, 4.5, 6.8}
	productNames[2] = "pencil"
	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[2])

	//slice usage 1:3 -> 1 is included 3 is excluded
	//slice usage :3 -> start beginning of the array 3 is excluded
	//slice usage 1: -> 1 is included at the end of the array

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
	featuredPrices = prices[1:]
	highlightedPrices := featuredPrices[:1]

	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]

	fmt.Println(highlightedPrices)
	//cap looks at the right side of the array that you slice
	// cap (capacity) means maximum num of element of the array, it is needed for memory allocation
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	testPrices := []float64{1.2, 2.3}
	testPrices[1] = 45.99
	//add a new value for slices use append builtin func it returns a new array rather than updating existing
	updatedPrices := append(testPrices, 5.99)

	fmt.Println(updatedPrices, testPrices)
	// or you can reassign existing slice again
	testPrices = append(testPrices, 5.99)

	fmt.Println(testPrices)

	discountPrices := []float64{1.23, 2.36}
	// merging 2 arrays you can use append func with 3 dots...
	testPrices = append(testPrices, discountPrices...)

	fmt.Println(testPrices)

}
