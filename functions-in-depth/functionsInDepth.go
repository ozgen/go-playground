package main

import "fmt"

type transFn func(int) int

/*
*
functions as value
anonymous functions
closures
recursion
variadic functions
*/
func main() {
	nums := []int{1, 2, 3, 4, 5}
	//doubled := doubleNumbers(&nums)
	doubled := transformNumbers(&nums, double)
	tripled := anotherTransformNumbers(&nums, triple)
	//this is the anonymous func definition
	transformed := transformNumbers(&nums, func(i int) int {
		return i * 4
	})

	factored := transformNumbers(&nums, createTransformFn(5))

	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(transformed)
	fmt.Println(factored)

	transFunc1 := getTransFunc(1)
	transFunc2 := getTransFunc(2)
	doubled = transformNumbers(&nums, transFunc1)
	tripled = transformNumbers(&nums, transFunc2)

	fmt.Println(doubled)
	fmt.Println(tripled)

	res := factorial(5)
	fmt.Println(res)

	sum := sumUp(1, 10, 15, -5)
	fmt.Println(sum)
	//or
	anotherSum := sumUp(nums...)
	fmt.Println(anotherSum)
}

func doubleNumbers(nums *[]int) []int {

	dNums := []int{}
	for _, val := range *nums {
		//dNums = append(dNums, val*2)
		dNums = append(dNums, double(val))
	}
	return dNums
}

func transformNumbers(nums *[]int, transform func(int) int) []int {

	dNums := []int{}
	for _, val := range *nums {
		dNums = append(dNums, transform(val))
	}
	return dNums
}
func anotherTransformNumbers(nums *[]int, fn transFn) []int {

	dNums := []int{}
	for _, val := range *nums {
		dNums = append(dNums, fn(val))
	}
	return dNums
}

// you can also define like that
func getTransFunc(selection int) transFn {
	if selection == 1 {
		return double
	} else {
		return triple
	}
}

/*
*
function creates a func with params
every anonymous func is closure
*/
func createTransformFn(factor int) func(int) int {
	// here is the anonymous func
	return func(i int) int {
		return i * factor
	}
}

func double(num int) int {
	return num * 2
}
func triple(num int) int {
	return num * 3
}

/*
*
example recursion implementation for factorial calculation
*/
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

/*
*
example variadic implementation for sumUp calculation
*/
func sumUp(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
