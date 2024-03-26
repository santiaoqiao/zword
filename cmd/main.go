package main

import (
	"fmt"
	"santiaoqiao.com/zword"
	"santiaoqiao.com/zword/pkg/docx"
)

func main() {

	doc, err := zword.OpenDocxFile("../tmp/aaa.docx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, c := range doc.Document.Body.Children {
		if c.TypeName() == docx.BodyTypeParagraph {
			p, ok := c.(*docx.Paragraph)
			if ok {
				fmt.Println(p.String())
			}
		}
	}
}
