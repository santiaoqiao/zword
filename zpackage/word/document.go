package word

import (
	"santiaoqiao.com/zoffice/zpackage/word/document"
)

type Document struct {
	Body Body `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
}

type Body struct {
	Children []document.Paragraph `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main p"`
}
