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
	Content   string    `json:"matter"`
	CreatedAt time.Time `json:"timestamp_for_creation"`
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Title: %v\n", note.Title)
	fmt.Printf("Content: %v\n", note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
