package main

import (
	"fmt"
	"santiaoqiao.com/zword"
)

func main() {

	_, err := zword.OpenDocxFile("../tmp/aaa.docx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println(doc.Body)
	//for _, c := range doc.Document.Body.Children {
	//	if c.TypeName() == docx.BodyTypeParagraph {
	//		p, ok := c.(*docx.Paragraph)
	//		if ok {
	//			fmt.Println(p.String())
	//		}
	//	}
	//}
}
