package main

import (
	"bufio"
	"fmt"
	"gokulskannan1993/note/note"
	"os"
	"strings"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving Note Failed!")
		return
	}

	fmt.Println("Note Saved Successfully")

}

func getNoteData() (string, string) {
	title := getUserInput(("Note title: "))

	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
