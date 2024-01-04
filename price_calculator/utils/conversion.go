package utils

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for idx, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			fmt.Println("Converting Price to float failed")
			fmt.Println(err)
			return nil, errors.New("Failed to convert strings to floats")
		}
		floats[idx] = floatValue
	}

	return floats, nil
}
