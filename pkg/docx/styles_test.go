package docx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStylesXml(t *testing.T) {
	doc, err := OpenDocxFile("../../test/aaa.docx")
	assert.NoError(t, err)
	// 读取DocDefaults
	assert.Equal(t, "Times New Roman", doc.Styles.DocDefaults.RPrDefault.fonts.Ascii)
	assert.Equal(t, 260, doc.Styles.LatentStyles.Count)
	assert.Equalf(t, 248, len(doc.Styles.LatentStyles.LsdExceptions), "the lenth of lsdExceptions is %d", len(doc.Styles.LatentStyles.LsdExceptions))
	assert.Equal(t, "Normal", doc.Styles.LatentStyles.LsdExceptions[0].Name)
	assert.Equal(t, "Colorful Grid Accent 6", doc.Styles.LatentStyles.LsdExceptions[len(doc.Styles.LatentStyles.LsdExceptions)-1].Name)
	assert.Equal(t, 14, len(doc.Styles.StyleSheets))
	assert.Equal(t, 24, *(doc.Styles.StyleSheets["1"].RPr.sizeCs))
	assert.Equal(t, "我的字体", doc.Styles.StyleSheets["14"].Name)
}
