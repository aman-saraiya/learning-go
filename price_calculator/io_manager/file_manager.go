package io_manager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFile  string
	OutputFile string
}

func NewFileManager(inputFile, outputFile string) FileManager {
	return FileManager{
		inputFile,
		outputFile,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFile)

	if err != nil {

		return nil, errors.New("Couldn't open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New("Reading File context failed")
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFile)
	if err != nil {
		return errors.New("Failed to create file.")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Failed to convert data to JSON.")
	}
	return nil
}
