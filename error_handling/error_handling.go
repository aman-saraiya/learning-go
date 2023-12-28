package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeIntToFile(value int) {
	balanceText := fmt.Sprint(value)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func readIntFromFile() (int64, error) {
	bytes, err := os.ReadFile("balance.txt")
	var balanceText int64
	if err != nil {
		balanceText = 10000
	} else {
		balanceText, err = strconv.ParseInt(string(bytes), 0, 0)
		if err != nil {
			balanceText = 10000
		}
	}
	fmt.Println("Balance text read from file:", balanceText)
	if err != nil {
		return balanceText, errors.New("Read or Parse Failed. Using default Value.")
	}
	return balanceText, nil
}
