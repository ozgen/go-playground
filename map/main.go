package main

import "fmt"

func main() {
	var colorMap map[int]string
	anotherColorMap := make(map[int]string)
	colors := map[int]string{
		1: "red",
		2: "green",
	}

	colors[3] = "white"
	delete(colors, 1)

	fmt.Println(colors)
	fmt.Println(colorMap)
	fmt.Println(anotherColorMap)
	printMap(colors)
}

func printMap(c map[int]string) {
	for key, value := range c {
		fmt.Println("key : ", key, " value: ", value)

	}

}
