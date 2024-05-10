package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	result := add(1, 2.4)
	printSomething(result)
	result2 := add("gege ", "hello")
	printSomething(result2)
	printSomething("Hello")
	title, content := getNoteData()
	todoText := getTodoData()

	todoNote, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todoNote)

	if err != nil {
		fmt.Println("Saving the todo failed.")
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

	fmt.Println("Saving the note succeeded!")

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println(value)
	}

	typedVal, ok := value.(int)
	if ok {
		fmt.Println("Integer + 1: ", typedVal+1)
	}
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content

}

func getTodoData() string {
	todoText := getUserInput("Todo Text:")

	return todoText

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
