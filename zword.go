package zword

import (
	"github.com/santiaoqiao/zword/pkg/docx"
)

func OpenDocxFile(filename string) (*docx.Docx, error) {
	return docx.OpenDocxFile(filename)
}
