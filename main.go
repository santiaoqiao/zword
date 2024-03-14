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

	for _, p := range docx.Document.Body.Children {
		fmt.Printf("%v", p)
	}
}
