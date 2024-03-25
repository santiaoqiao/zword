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
	//fmt.Println(doc.Body)
	for _, c := range doc.Body.Children {
		if c.TypeName() == "p" {
			p, ok := c.(*docx.Paragraph)
			if ok {
				fmt.Println(p.String() + "...")
			}
		}
	}
}
