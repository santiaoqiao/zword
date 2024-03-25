package docx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStylesXml(t *testing.T) {
	doc, err := OpenDocxFile("../../test/aaa.docx")
	assert.NoError(t, err)
	// 读取DocDefaults
	assert.Equal(t, "Times New Roman", doc.Styles.DocDefaults.RPrDefault.Fonts.Ascii)
}
