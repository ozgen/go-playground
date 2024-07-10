package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var rate = 5.5
	var years float64

	printText("Investment Amount : ")
	fmt.Scan(&investmentAmount)

	printText("Return rate : ")
	fmt.Scan(&rate)

	printText("Years : ")
	fmt.Scan(&years)

	fv, rfv := calculateFutureAmounts(investmentAmount, rate, years)

	fmt.Println(fv)
	fmt.Println(rfv)

	//iAmount := 1000.0
	//rate2 := 5.5
	//years2 := 10.0
	//
	//
	//retVal2 := iAmount * math.Pow(1+rate2/100, years2)
	//fmt.Println(retVal2)
	//realVal := retVal2 / math.Pow(1+inflationRate/100, years2)

	//fmt.Println(realVal)
}

func printText(text string) {
	fmt.Print(text)
}

func calculateFutureAmounts(investmentAmount float64, rate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+rate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
