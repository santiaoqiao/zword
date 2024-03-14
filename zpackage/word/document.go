package word

import "santiaoqiao.com/zoffice/zpackage/word/document"

type Document struct {
	Body document.Body `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
}
