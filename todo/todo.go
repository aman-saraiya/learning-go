package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"content"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Todo creation failed. Content cannot be blank.")
	}
	todo := Todo{
		content,
	}
	return todo, nil
}

func (todo Todo) Display() {
	fmt.Println("Content:", todo.Content)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
