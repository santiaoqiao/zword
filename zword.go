package zword

import (
	"santiaoqiao.com/zword/pkg/docx"
)

func OpenDocxFile(filename string) (*docx.Docx, error) {
	return docx.OpenDocxFile(filename)
}
