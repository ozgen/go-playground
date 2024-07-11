package main

import (
	"fmt"
	"priceCalculator/filemanager"
	"priceCalculator/prices"
)

func main() {

	//prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	//result := make(map[float64][]float64)

	//for _, taxRate := range taxRates {
	//	taxIncludedPrices := make([]float64, len(prices))
	//	for indexOfPrice, price := range prices {
	//		taxIncludedPrices[indexOfPrice] = price * (1 + taxRate)
	//	}
	//	result[taxRate] = taxIncludedPrices
	//}

	//fmt.Println(result)

	for _, taxRate := range taxRates {
		fileManager := filemanager.New(prices.Prices_File, fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdManager := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		job.Process()
	}
}
