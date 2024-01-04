package prices

import (
	"fmt"

	"github.com/aman-saraiya/learning-go/price_calculator/io_manager"
	"github.com/aman-saraiya/learning-go/price_calculator/utils"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64              `json:"tax_rate"`
	InputPrices       []float64            `json:"input_prices"`
	TaxIncludedPrices map[string]float64   `json:"tax_included_prices"`
	IOManager         io_manager.IOManager `json:"-"`
}

func NewTaxIncludedPriceJob(taxRate float64, fm io_manager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10.12, 15.10, 110.12, 15.17},
		IOManager:   fm,
	}
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	job.TaxIncludedPrices = make(map[string]float64, len(job.InputPrices))
	for _, price := range job.InputPrices {
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	return job.IOManager.WriteResult(job)
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	prices, err := utils.StringsToFloats(lines)
	if err != nil {
		return err
	}
	job.InputPrices = prices
	return nil
}
