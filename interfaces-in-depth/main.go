package main

import (
	"bufio"
	"fmt"
	"interfaces_in_depths/note"
	"interfaces_in_depths/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func main() {

	printAnything(1)
	printAnything(1.5)
	printAnything("hello test")
	generic := addGeneric(1, 2)
	printAnything(generic)
	title, content := getNoteData()

	todoText := getUserInput("Add todo:")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}
}

// interface{} : allows to accept any kind of struct
func printAnything(data interface{}) {
	//fmt.Println(data)
	switch data.(type) {
	case int:
		fmt.Println("Integer: ", data)
	case float64:
		fmt.Println("float64: ", data)
	case string:
		fmt.Println("string: ", data)
	}
	// or
	intVal, ok := data.(int)
	if ok {
		fmt.Println("int value: ", intVal)
	}
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}

	fmt.Println("Saving the data succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

/*
*
generic types
*/
func addGeneric[T int | float64 | string](a, b T) T {
	return a + b
}
