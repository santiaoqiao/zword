package main

import (
	"fmt"
	"github.com/santiaoqiao/zword"
	"github.com/santiaoqiao/zword/pkg/docx"
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
				for _, c1 := range p.Children {
					if c1.TypeName() == docx.ParagraphTypeRun {
						r, _ := c1.(*docx.Run)
						if r.RunProperty.Bold() != nil {
							fmt.Printf("%s(%v)\n", r.String(), *r.RunProperty.Bold())
						} else {
							fmt.Printf("%s(nil)\n", r.String())
						}

					}
				}
			}
		}
	}
}
