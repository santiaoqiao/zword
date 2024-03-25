package main

import (
	"fmt"
	"santiaoqiao.com/zword"
	"santiaoqiao.com/zword/internal/xmldocx/word/stroies"
)

func main() {

	doc, err := zword.OpenDocxFile("../tmp/aaa.xmldocx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println(doc.Body)
	for _, c := range doc.Body.Children {
		if c.TypeName() == "p" {
			p, ok := c.(*stroies.Paragraph)
			if ok {
				fmt.Println(p.String() + "...")
			}
		}
	}
}
