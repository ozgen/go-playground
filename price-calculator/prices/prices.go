package prices

import (
	"fmt"
	"priceCalculator/conversion"
	"priceCalculator/iomanager"
)

const Prices_File = "prices.txt"

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"io_manager"` // if you want to ignore the value use `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

// we need to set as pointer because we want to update existing jobs inputPrices
func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("Could not read the file!")
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("value could not be parsed to float64")
		fmt.Println(err)
		return
	}
	job.InputPrices = prices
}

func (job TaxIncludedPriceJob) Process() {

	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
	job.TaxIncludedPrices = result
	job.IOManager.WriteJSON(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: iom,
	}
}
