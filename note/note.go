package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// here the `json:"<name>"` tags are used
// by the json.Marshall function to rename the fields
// before saving the struct as json to file
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// we could have returned a Note pointer here but it is preferable to not
// use a pointer because of the extra overhead that takes
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Note creation failed. Title and content cannot be blank.")
		// Note we cannot return nil as a substitute for struct
		// nil only works with pointer, channel, func, interface, map, or slice type.
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

	// NOTE: fields of the passed struct that doesn't
	// start with capital letters are not accessible in other packages
	// so json.Marshall only will add the fields that are accessible
	data, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
