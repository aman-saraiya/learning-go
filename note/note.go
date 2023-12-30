package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Note creation failed. Title and content cannot be blank.")

	}
	note := Note{
		title,
		content,
		time.Now(),
	}
	return note, nil
}

func (note Note) Display() {
	fmt.Println("Title:", note.Title)
	fmt.Println("Content:", note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	data, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
