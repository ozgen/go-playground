package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// write uppercase the function name to export
func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
