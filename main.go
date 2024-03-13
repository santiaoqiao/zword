package main

import (
	"fmt"
)

func main() {
	docx := &Docx{}
	err := docx.Read("./tmp/官僚资本主义.docx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, k := range docx.ContentTypes.DefaultItems {
		fmt.Println(k)
	}

	for _, k := range docx.ContentTypes.OverrideItem {
		fmt.Println(k)
	}

	for _, p := range docx.Document.Body.Children {
		fmt.Println(p)
	}
	//fmt.Println(*docx)
}
