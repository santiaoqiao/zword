package main

import (
	"fmt"
)

func main() {
	docx := &Docx{}
	err := docx.Read("./tmp/aaa.docx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(&docx.Document.Body)
}
