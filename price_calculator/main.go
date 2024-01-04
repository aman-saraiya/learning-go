package main

import (
	"fmt"

	"github.com/aman-saraiya/learning-go/price_calculator/io_manager"
	"github.com/aman-saraiya/learning-go/price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.15, 0.2}

	for _, taxRate := range taxRates {
		fm := io_manager.NewFileManager("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		err := priceJob.Process()
		if err != nil {
			fmt.Printf("Failed to process job for taxRate %f\n", taxRate)
		}
		// cmd_m := io_manager.NewCMDManager()
		// priceJob := prices.NewTaxIncludedPriceJob(taxRate, cmd_m)
		// priceJob.Process()
	}
}
