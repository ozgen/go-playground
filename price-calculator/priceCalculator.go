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
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	for i, taxRate := range taxRates {
		fileManager := filemanager.New(prices.Prices_File, fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdManager := cmdmanager.New()
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		job := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		go job.Process(doneChans[i], errorChans[i])
	}

	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			fmt.Println(err)
		case <-doneChans[i]:
			fmt.Println("Done!")
		}
	}
}
