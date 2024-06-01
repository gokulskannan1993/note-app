package main

import (
	"fmt"

	"gokulskannan1993/note/note"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.Note(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func getNoteData() (string, string) {
	title := getUserInput(("Note title"))

	content := getUserInput("Note Content")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string

	fmt.Scanln(&value)

	return value
}
