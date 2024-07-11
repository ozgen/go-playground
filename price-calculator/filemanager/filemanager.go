package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string `json:"input_file_path"`
	OutputFilePath string `json:"output_file_path"`
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (manager FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(manager.InputFilePath)
	if err != nil {
		fmt.Println("File can not be opened!")
		fmt.Println(err)
		//file.Close()
		return nil, err
	}

	// we can use defer command to close the file in correct time rather than adding close method to all edge cases
	// because it triggers when the outer function is DONE.
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		//file.Close()
		return nil, err
	}
	//file.Close()
	return lines, nil
}

// data could be any type
func (manager FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(manager.OutputFilePath)
	if err != nil {
		//file.Close()
		return err
	}

	// we can use defer command to close the file in correct time rather than adding close method to all edge cases
	// because it triggers when the outer function is DONE.
	defer file.Close()

	time.Sleep(3 * time.Second) // simulate a slow, long-taking task

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		//file.Close()
		return err
	}
	//file.Close()
	return nil
}
