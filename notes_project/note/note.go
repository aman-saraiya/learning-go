package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
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
	fmt.Println("Title:", note.title)
	fmt.Println("Content:", note.content)
}
