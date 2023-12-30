package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aman-saraiya/learning-go/interfaces/note"
	"github.com/aman-saraiya/learning-go/interfaces/todo"
)

func main() {
	title, content := getNoteData()
	myNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	myNote.Display()
	// duplicate code that could be removed by using interfaces
	/*
		err = myNote.Save()
		if err != nil {
			fmt.Println("Saving the Note failed.")
		} else {
			fmt.Println("Note saved successfully to file.")
		}
	*/

	// duplication removed using interfaces
	saveData(myNote)

	content = getTodoData()
	myTodo, err := todo.New(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	myTodo.Display()
	// duplicate code that could be removed by using interfaces
	/*
		err = myTodo.Save()
		if err != nil {
			fmt.Println("Saving the Todo failed.")
		} else {
			fmt.Println("Todo saved successfully to file.")
		}
	*/

	//duplication removed using interfaces
	saveData(myTodo)

	//NOTE: We're able to pass myTodo and myNote in place where a
	// saver type is required because saver is an interface and it
	// just enforces that a valid value for it must implement
	// Save() method. And both Todo and Note have Save method.
}

// we cannot specify note.Note or todo.Todo as input param type if we want
// to use this function with both types. We need interfaces.
func saveData(dataStruct saver) {
	err := dataStruct.Save()
	if err != nil {
		fmt.Println("Saving the data struct failed.")
	} else {
		fmt.Println("Data struct saved successfully to file.")
	}
}

// creating interface
type saver interface {
	Save() error
}

// side note, there's a convention in Go that if we define an interface with
// only one method, than our interface name is method_name + "er", it's not a
// must do but a common convention

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")

	// this is required as in windows new line is \r\n
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Enter Note title: ")
	content := getUserInput("Enter Note content: ")
	return title, content
}

func getTodoData() string {
	content := getUserInput("Enter Todo content: ")
	return content
}
