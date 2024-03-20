package main

import (
	"fmt"
	docx2 "santiaoqiao.com/zword/internal/docx"
)

func main() {
	docx := &docx2.Docx{}
	err := docx.Read("../tmp/aaa.docx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(&docx.Document.Body)
}