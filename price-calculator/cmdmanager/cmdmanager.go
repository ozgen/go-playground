package cmdmanager

import (
	"fmt"
)

type CMDManager struct {
}

func New() CMDManager {
	return CMDManager{}
}

func (manager CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Add price and confirm price with ENTER")
	var lines []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		lines = append(lines, price)

	}
	return lines, nil
}

func (manager CMDManager) WriteJSON(data interface{}) error {
	fmt.Println(data)
	return nil
}
