package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// importing custom packages or third party packages
	// module_path/package
	"github.com/aman-saraiya/learning-go/interfaces/note"
)

func main() {
	title, content := getNoteData()
	myNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	myNote.Display()
	err = myNote.Save()
	if err != nil {
		fmt.Println("Saving the Note failed.")
	} else {
		fmt.Println("Note saved successfully to file.")
	}
}

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
	title := getUserInput("Enter title: ")
	content := getUserInput("Enter content: ")
	return title, content
}
